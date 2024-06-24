package server

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"

	"github.com/gofiber/contrib/websocket"
	"google.golang.org/protobuf/proto"

	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/ws"
)

func (s *Server) handlerWS(conn *websocket.Conn) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	log := logger.FromContext(ctx)

	log.Debugf("New connection from: %s", conn.RemoteAddr().String())

	cancel := sync.OnceFunc(func() {
		ctxCancel()
		ws.WriteCloseMessage(conn)
		log.Debugf("Connection closed from: %s", conn.RemoteAddr().String())
	})
	defer cancel()

	token := conn.Query("token")
	if token == "" {
		log.Warnf("Missing token")
		return
	}

	client, ok := s.clients.Get(token)
	if !ok {
		log.Errorf("Unknown token: %s", token)
		return
	}

	connectionID := rand.Uint64()
	log = log.With("player-id", client.PlayerID, "connection-id", connectionID)
	ctx = logger.ToContext(ctx, log)

	responses := make(chan *socket.Response, 32)
	state := &ClientState{
		Context:      ctx,
		ConnectionID: connectionID,
		Responses:    responses,
	}

	client.StateLock.Lock()
	if client.State == nil {
		client.State = state
	} else {
		client.StateLock.Unlock()
		log.Errorf("Player already connected")
		return
	}
	client.StateLock.Unlock()

	defer func() {
		client.StateLock.Lock()
		client.State = nil
		client.StateLock.Unlock()
	}()

	state.Responses <- &socket.Response{
		Type: &socket.Response_Info_{
			Info: &socket.Response_Info{
				PlayerId: client.PlayerID,
			},
		},
	}

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-responses:
				if !ok {
					return
				}

				if err := msg.ValidateAll(); err != nil {
					log.Errorf("Invalid response: %s", err)
					return
				}

				data, err := proto.Marshal(msg)
				if err != nil {
					log.Errorf("Marshaling response: %s", err)
					return
				}

				if !ws.WriteMessage(log, conn, data) {
					return
				}
			}
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			// Continue
		}

		data, ok := ws.ReadMessage(log, conn)
		if !ok {
			return
		}

		msg := &socket.Request{}
		if err := proto.Unmarshal(data, msg); err != nil {
			log.Errorf("Unmarshaling request: %s", err)
			return
		}

		if err := msg.ValidateAll(); err != nil {
			log.Errorf("Invalid request: %s", err)
			continue
		}

		if err := s.processRequest(client, msg); err != nil {
			log.Errorf("Processing request: %s", err)
			continue
		}
	}
}

func (s *Server) processRequest(client *Client, request *socket.Request) error {
	switch req := request.Type.(type) {
	case *socket.Request_PlayerMove:
		msg := &socket.Response{
			Type: &socket.Response_PlayerMove_{
				PlayerMove: &socket.Response_PlayerMove{
					PlayerId: client.PlayerID,
					Pos:      req.PlayerMove,
				},
			},
		}

		s.clients.ForEach(func(_ string, otherClient *Client) bool {
			otherState := otherClient.SafeState()
			if otherState == nil || otherState.ConnectionID == client.State.ConnectionID {
				return true
			}

			otherState.SendMessage(msg)
			return true
		})
	default:
		client.State.SendError(&socket.Response_Error{
			Code:    socket.Response_Error_UNSUPPORTED_REQUEST,
			Message: fmt.Sprintf("%T", request.Type),
		})
		logger.FromContext(client.State).Errorf("Unsupported request: %T", request.Type)
		return nil
	}
	return nil
}

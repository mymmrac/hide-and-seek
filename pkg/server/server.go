package server

import (
	"context"
	"errors"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/contrib/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/api"
	"github.com/mymmrac/hide-and-seek/pkg/logger"
)

type Server struct {
	playerLock sync.RWMutex
	players    map[uint64]*Client
}

type Client struct {
	ConnWrite chan *api.Msg
}

func NewServer() *Server {
	return &Server{
		playerLock: sync.RWMutex{},
		players:    make(map[uint64]*Client),
	}
}

func (s *Server) Handler(conn *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
	log := logger.FromContext(ctx)

	log.Debugf("New connection from: %s", conn.RemoteAddr().String())

	cancel = sync.OnceFunc(func() {
		cancel()

		if err := conn.WriteControl(websocket.CloseMessage, nil, time.Now().Add(time.Second)); err != nil {
			log.Errorf("Error sending close message: %s", err)
		}

		if err := conn.Close(); err != nil {
			log.Errorf("Error closing connection: %s", err)
		}

		log.Debugf("Connection closed from: %s", conn.RemoteAddr().String())
	})
	defer cancel()

	connectionID, err := strconv.ParseUint(conn.Query("id"), 10, 64)
	if err != nil {
		log.Errorf("Error parsing connection ID: %s", err)
		return
	}

	log = log.With("connection-id", connectionID)
	ctx = logger.ToContext(ctx, log)

	client := &Client{
		ConnWrite: make(chan *api.Msg, 32),
	}

	s.playerLock.Lock()
	s.players[connectionID] = client
	s.playerLock.Unlock()

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-client.ConnWrite:
				if !ok {
					return
				}

				data, err := msg.Marshal()
				if err != nil {
					log.Errorf("Error marshaling message: %s", err)
					return
				}

				err = conn.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					if websocket.IsCloseError(err, websocket.CloseNormalClosure) || errors.Is(err, net.ErrClosed) {
						return
					}

					if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
						log.Errorf("Unexpected close: %s", err)
						return
					}

					log.Errorf("Error writing message: %s", err)
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

		msgType, data, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				return
			}

			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("Unexpected close: %s", err)
				return
			}

			log.Errorf("Error reading message: %s", err)
			return
		}
		if msgType != websocket.BinaryMessage {
			continue
		}

		msg := &api.Msg{}
		if err = msg.Unmarshal(data); err != nil {
			log.Errorf("Error unmarshaling message: %s", err)
			return
		}

		// log.Debugf("Received message: %+v", msg)

		s.playerLock.Lock()
		for id, player := range s.players {
			if id == connectionID {
				continue
			}

			select {
			case player.ConnWrite <- msg:
				// Sent
			default:
				log.Error("Write buffer is full")
			}
		}
		s.playerLock.Unlock()
	}
}

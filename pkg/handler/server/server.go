package server

import (
	"context"
	"strconv"
	"sync"

	"github.com/gofiber/contrib/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/api"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/ws"
)

type Server struct {
	playerLock sync.RWMutex
	players    map[uint64]*Client
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
		ws.WriteCloseMessage(conn)
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

	responses := make(chan *api.Msg, 32)
	client := &Client{
		Responses: responses,
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
			case msg, ok := <-responses:
				if !ok {
					return
				}

				data, err := msg.Marshal()
				if err != nil {
					log.Errorf("Error marshaling message: %s", err)
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

		msg := &api.Msg{}
		if err = msg.Unmarshal(data); err != nil {
			log.Errorf("Error unmarshaling message: %s", err)
			return
		}

		s.playerLock.Lock()
		for id, player := range s.players {
			if id == connectionID {
				continue
			}

			select {
			case player.Responses <- msg:
				// Sent
			default:
				log.Error("Write buffer is full")
			}
		}
		s.playerLock.Unlock()
	}
}

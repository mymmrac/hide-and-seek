package server

import (
	"context"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/api"
	"github.com/mymmrac/hide-and-seek/pkg/module/collection"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

type Server struct {
	clients   *collection.SyncMap[string, *Client] // Token -> Client
	broadcast chan *socket.Response
}

func NewServer() *Server {
	return &Server{
		clients:   collection.NewSyncMap[string, *Client](),
		broadcast: make(chan *socket.Response, 32),
	}
}

func (s *Server) RegisterHandlers(router fiber.Router) {
	router.Post("/start", api.ProtoHandler(s.handlerStart))
	router.Get("/ws", websocket.New(s.handlerWS))
}

func (s *Server) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg, ok := <-s.broadcast:
			if !ok {
				return
			}

			s.clients.ForEach(func(_ string, client *Client) bool {
				state := client.SafeState()
				if state == nil {
					return true
				}
				state.SendMessage(msg)
				return true
			})
		}
	}
}

func (s *Server) BroadcastMessage(ctx context.Context, msg *socket.Response) {
	select {
	case s.broadcast <- msg:
		// Sent
	default:
		logger.FromContext(ctx).Error("Broadcast buffer is full")
	}
}

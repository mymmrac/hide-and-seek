package server

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"github.com/mymmrac/hide-and-seek/pkg/module/api"
	"github.com/mymmrac/hide-and-seek/pkg/module/collection"
)

type Server struct {
	clients *collection.SyncMap[string, *Client] // Token -> Client
}

func NewServer() *Server {
	return &Server{
		clients: collection.NewSyncMap[string, *Client](),
	}
}

func (s *Server) RegisterHandlers(router fiber.Router) {
	router.Post("/start", api.ProtoHandler(s.handlerStart))
	router.Get("/ws", websocket.New(s.handlerWS))
}

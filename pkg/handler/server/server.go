package server

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"

	"github.com/mymmrac/hide-and-seek/pkg/module/api"
)

type Server struct {
	playerLock sync.RWMutex // TODO: Use channel instead
	players    map[string]*Client
}

func NewServer() *Server {
	return &Server{
		playerLock: sync.RWMutex{},
		players:    make(map[string]*Client),
	}
}

func (s *Server) RegisterHandlers(router fiber.Router) {
	router.Post("/start", api.ProtoHandler(s.handlerStart))
	router.Get("/ws", websocket.New(s.handlerWS))
}

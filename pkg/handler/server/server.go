package server

import "sync"

type Server struct {
	playerLock sync.RWMutex // TODO: Use channel instead
	players    map[uint64]*Client
}

func NewServer() *Server {
	return &Server{
		playerLock: sync.RWMutex{},
		players:    make(map[uint64]*Client),
	}
}

package server

import (
	"context"
	"sync"

	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

type Client struct {
	PlayerID   uint64
	PlayerName string

	StateLock sync.RWMutex
	State     *ClientState
}

func (c *Client) SafeState() *ClientState {
	c.StateLock.RLock()
	defer c.StateLock.RUnlock()
	return c.State
}

type ClientState struct {
	context.Context

	ConnectionID uint64
	Responses    chan<- *socket.Response
}

func (s *ClientState) SendMessage(msg *socket.Response) {
	select {
	case s.Responses <- msg:
		// Sent
	default:
		logger.FromContext(s).Error("Write buffer is full")
	}
}

func (s *ClientState) SendError(err *socket.Response_Error) {
	s.SendMessage(&socket.Response{
		Type: &socket.Response_Error_{
			Error: err,
		},
	})
}

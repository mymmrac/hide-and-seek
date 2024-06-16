package server

import (
	"context"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/gofiber/contrib/websocket"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Handler(conn *websocket.Conn) {
	ctx, cancel := context.WithCancel(context.Background())
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

		log.Debugf("Received message: %s", string(data))
	}
}

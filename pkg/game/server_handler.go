package game

import (
	"context"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/charmbracelet/log"
	"github.com/fasthttp/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/api"
)

func (g *Game) HandleConnection(conn *websocket.Conn) {
	ctx, ctxCancel := context.WithCancel(g.ctx)
	log.Debugf("Connected to server: %s", conn.RemoteAddr().String())

	cancel := sync.OnceFunc(func() {
		ctxCancel()

		if err := conn.WriteControl(websocket.CloseMessage, nil, time.Now().Add(time.Second)); err != nil {
			log.Errorf("Error sending close message: %s", err)
		}

		if err := conn.Close(); err != nil {
			log.Errorf("Error closing connection: %s", err)
		}

		log.Debugf("Connection to server closed: %s", conn.RemoteAddr().String())
		g.events <- EventDisconnectedFromServer
		g.wg.Done()
	})
	defer cancel()

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-g.connWrite:
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
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) || errors.Is(err, net.ErrClosed) {
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
		select {
		case g.connRead <- msg:
		default:
			log.Errorf("Connection read buffer full")
		}
	}
}

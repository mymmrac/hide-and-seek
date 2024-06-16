package game

import (
	"context"
	"sync"

	"github.com/fasthttp/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/api"
	"github.com/mymmrac/hide-and-seek/pkg/logger"
	"github.com/mymmrac/hide-and-seek/pkg/ws"
)

func (g *Game) handleConnection(conn *websocket.Conn) {
	ctx, cancel := context.WithCancel(g.ctx)
	log := logger.FromContext(ctx)

	log.Debugf("Connected to server: %s", conn.RemoteAddr().String())

	cancel = sync.OnceFunc(func() {
		cancel()
		ws.WriteCloseMessage(conn)
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
		if err := msg.Unmarshal(data); err != nil {
			log.Errorf("Error unmarshaling message: %s", err)
			return
		}

		select {
		case g.connRead <- msg:
			// Process message
		default:
			log.Errorf("Connection read buffer full")
		}
	}
}

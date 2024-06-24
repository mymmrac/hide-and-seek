package game

import (
	"context"
	"sync"

	"github.com/fasthttp/websocket"
	"google.golang.org/protobuf/proto"

	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/ws"
)

func (g *Game) handleConnection(conn *websocket.Conn) {
	ctx, ctxCancel := context.WithCancel(g.ctx)
	log := logger.FromContext(ctx)

	log.Debugf("Connected to server: %s", conn.RemoteAddr().String())

	cancel := sync.OnceFunc(func() {
		ctxCancel()
		ws.WriteCloseMessage(conn)
		log.Debugf("Connection to server closed: %s", conn.RemoteAddr().String())
		g.events <- EventDisconnectedFromServer
		g.wg.Done() // WG: Server connection
	})
	defer cancel()

	go func() {
		defer cancel()

		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-g.requests:
				if !ok {
					return
				}

				if err := msg.ValidateAll(); err != nil {
					log.Errorf("Invalid request: %s", err)
					continue
				}

				data, err := proto.Marshal(msg)
				if err != nil {
					log.Errorf("Marshaling request: %s", err)
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

		msg := &socket.Response{}
		if err := proto.Unmarshal(data, msg); err != nil {
			log.Errorf("Unmarshaling response: %s", err)
			return
		}

		if err := msg.ValidateAll(); err != nil {
			log.Errorf("Invalid response: %s", err)
			continue
		}

		select {
		case g.responses <- msg:
			// Process message
		default:
			log.Errorf("Connection read buffer full")
		}
	}
}

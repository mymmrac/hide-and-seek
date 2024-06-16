package game

import (
	"strconv"

	"github.com/fasthttp/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/api"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

func (g *Game) connectToServer() {
	conn, _, err := websocket.DefaultDialer.DialContext(
		g.ctx,
		"ws://localhost:4242?id="+strconv.FormatUint(g.connectionID, 10),
		nil,
	)
	if err != nil {
		logger.FromContext(g.ctx).Errorf("Error connecting to server: %s", err)
		g.events <- EventDisconnectedFromServer
		return
	}

	g.requests = make(chan *api.Msg, 32)
	g.responses = make(chan *api.Msg, 32)

	g.wg.Add(1) // WG: Server connection
	go g.handleConnection(conn)

	g.events <- EventConnectedToServer
}

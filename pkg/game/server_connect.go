package game

import (
	"strconv"

	"github.com/fasthttp/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/api"
	"github.com/mymmrac/hide-and-seek/pkg/logger"
)

func (g *Game) ConnectToServer() {
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

	g.connWrite = make(chan *api.Msg, 32)
	g.connRead = make(chan *api.Msg, 32)

	g.wg.Add(1)
	go g.HandleConnection(conn)

	g.events <- EventConnectedToServer
}

package game

import (
	"strconv"

	"github.com/charmbracelet/log"
	"github.com/fasthttp/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/api"
)

func (g *Game) ConnectToServer() {
	conn, _, err := websocket.DefaultDialer.DialContext(
		g.ctx,
		"ws://localhost:4242?id="+strconv.FormatUint(g.connectionID, 10),
		nil,
	)
	if err != nil {
		log.Errorf("Error connecting to server: %s", err)
		g.events <- EventDisconnectedFromServer
		return
	}

	g.connWrite = make(chan *api.Msg, 32)
	g.connRead = make(chan *api.Msg, 32)

	g.wg.Add(1)
	go g.HandleConnection(conn)

	g.events <- EventConnectedToServer
}

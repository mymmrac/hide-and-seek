package game

import (
	"fmt"
	"math/rand/v2"
	"strconv"

	"github.com/fasthttp/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/api/communication"
	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/api"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

const serverHostname = "localhost:4242"

func (g *Game) connectToServer() {
	resp, err := api.ProtoCall[communication.Start_Response](
		g.httpClient,
		fmt.Sprintf("http://%s/start", serverHostname),
		&communication.Start_Request{
			Username: "test" + strconv.FormatUint(rand.Uint64(), 10),
		},
	)
	if err != nil {
		logger.FromContext(g.ctx).Errorf("Error starting game session: %s", err)
		g.events <- EventDisconnectedFromServer
		return
	}

	conn, _, err := websocket.DefaultDialer.DialContext(
		g.ctx,
		fmt.Sprintf("ws://%s/ws?token=", serverHostname)+resp.GetResult().Token,
		nil,
	)
	if err != nil {
		logger.FromContext(g.ctx).Errorf("Error connecting to server: %s", err)
		g.events <- EventDisconnectedFromServer
		return
	}

	g.requests = make(chan *socket.Request, 32)
	g.responses = make(chan *socket.Response, 32)

	g.wg.Add(1) // WG: Server connection
	go g.handleConnection(conn)

	g.events <- EventConnectedToServer
}

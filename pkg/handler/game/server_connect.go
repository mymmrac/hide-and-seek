package game

import (
	"fmt"

	"github.com/mymmrac/hide-and-seek/pkg/api/communication"
	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/api"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

func (g *Game) connectToServer() {
	resp, err := api.ProtoCall[communication.Start_Response](
		g.ctx,
		g.httpClient,
		fmt.Sprintf("http://%s/start", g.serverAddress),
		&communication.Start_Request{
			Username: g.player.Name,
		},
	)
	if err != nil {
		logger.FromContext(g.ctx).Errorf("Error starting game session: %s", err)
		g.events <- EventDisconnectedFromServer
		return
	}

	conn, err := g.httpClient.WS(g.ctx, fmt.Sprintf("ws://%s/ws?token=", g.serverAddress)+resp.GetResult().Token)
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

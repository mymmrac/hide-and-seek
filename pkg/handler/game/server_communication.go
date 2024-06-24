package game

import (
	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
)

func (g *Game) sendMessage(msg *socket.Request) {
	select {
	case g.requests <- msg:
		// Sent
	default:
		logger.FromContext(g.ctx).Errorf("Write buffer is full")
	}
}

func (g *Game) processMessages() {
	for {
		select {
		case msg, ok := <-g.responses:
			if !ok {
				return
			}
			g.processMessage(msg)
		default:
			return
		}
	}
}

func (g *Game) processMessage(msg *socket.Response) {
	switch resp := msg.Type.(type) {
	case *socket.Response_Bulk_:
		for _, response := range resp.Bulk.Responses {
			g.processMessage(response)
		}
	case *socket.Response_Error_:
		logger.FromContext(g.ctx).Errorf("Server error: %s:%s", resp.Error.Code.String(), resp.Error.Message)
	case *socket.Response_Info_:
		g.info = resp.Info
		logger.FromContext(g.ctx).Infof("Info: %+v", resp.Info)
	case *socket.Response_PlayerMove_:
		g.playerLock.Lock()
		g.players[resp.PlayerMove.PlayerId] = space.Vec2F{
			X: resp.PlayerMove.Pos.X,
			Y: resp.PlayerMove.Pos.Y,
		}
		g.playerLock.Unlock()
	default:
		logger.FromContext(g.ctx).Errorf("Unknown response type: %T", resp)
	}
}

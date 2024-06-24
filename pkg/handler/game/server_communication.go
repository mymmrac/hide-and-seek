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

		g.players.Lock()
		players := g.players.Raw()
		for _, player := range resp.Info.Players {
			players[player.Id] = &Player{
				Name: player.Username,
				Pos:  space.Vec2F{},
			}
		}
		delete(players, resp.Info.PlayerId)
		g.players.Unlock()

		logger.FromContext(g.ctx).Infof("Info: %+v", resp.Info)
	case *socket.Response_PlayerJoin_:
		if g.info != nil && g.info.PlayerId == resp.PlayerJoin.Id {
			return
		}

		g.players.Set(resp.PlayerJoin.Id, &Player{
			Name: resp.PlayerJoin.Username,
			Pos:  space.Vec2F{},
		})
		logger.FromContext(g.ctx).Infof("Player joined: %+v", resp.PlayerJoin)
	case *socket.Response_PlayerLeave:
		if g.info != nil && g.info.PlayerId == resp.PlayerLeave {
			return
		}

		g.players.Remove(resp.PlayerLeave)
		logger.FromContext(g.ctx).Infof("Player left: %+v", resp.PlayerLeave)
	case *socket.Response_PlayerMove_:
		player, ok := g.players.Get(resp.PlayerMove.PlayerId)
		if !ok {
			logger.FromContext(g.ctx).Errorf("Unknown player: %d", resp.PlayerMove.PlayerId)
			return
		}

		player.Pos = space.Vec2F{
			X: resp.PlayerMove.Pos.X,
			Y: resp.PlayerMove.Pos.Y,
		}
	default:
		logger.FromContext(g.ctx).Errorf("Unknown response type: %T", resp)
	}
}

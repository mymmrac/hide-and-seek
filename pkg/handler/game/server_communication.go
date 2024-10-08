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
		for _, playerJoin := range resp.Info.Players {
			if playerJoin.Id == resp.Info.PlayerId {
				continue
			}

			player := NewPlayer(g.cw)
			player.Name = playerJoin.Username

			players[playerJoin.Id] = player
		}
		g.players.Unlock()

		logger.FromContext(g.ctx).Infof("Info: %+v", resp.Info)
	case *socket.Response_PlayerJoin_:
		if g.info != nil && g.info.PlayerId == resp.PlayerJoin.Id {
			return
		}

		player := NewPlayer(g.cw)
		player.Name = resp.PlayerJoin.Username

		g.players.Set(resp.PlayerJoin.Id, player)

		logger.FromContext(g.ctx).Infof("Player joined: %+v", resp.PlayerJoin)
	case *socket.Response_PlayerLeave:
		if g.info != nil && g.info.PlayerId == resp.PlayerLeave {
			return
		}

		player, ok := g.players.GetAndRemove(resp.PlayerLeave)
		if ok {
			g.cw.Remove(player.Collider)
		}

		logger.FromContext(g.ctx).Infof("Player left: %+v", resp.PlayerLeave)
	case *socket.Response_PlayerState_:
		state := resp.PlayerState
		player, ok := g.players.Get(state.PlayerId)
		if !ok {
			logger.FromContext(g.ctx).Errorf("Unknown player: %d", state.PlayerId)
			return
		}

		oldPos := player.Collider.Position()
		newPos := space.Vec2F{
			X: state.Pos.X,
			Y: state.Pos.Y,
		}

		if newPos != oldPos {
			player.Collider.SetPosition(newPos)
			player.UpdatePosition()
		}

		player.Dir.X = int(state.Dir.X)
		player.Dir.Y = int(state.Dir.Y)
		player.Moving = state.Moving
	default:
		logger.FromContext(g.ctx).Errorf("Unknown response type: %T", resp)
	}
}

package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

func (g *Game) Update() error {
	select {
	case <-g.ctx.Done():
		g.wg.Done() // WG: Game loop
		return ebiten.Termination
	default:
		// Continue
	}

	if ebiten.IsWindowBeingClosed() || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.cancel()
		return nil
	}

	select {
	case event := <-g.events:
		switch event {
		case EventStartServer:
			go g.connectToServer()
		case EventConnectedToServer:
			logger.FromContext(g.ctx).Info("Connected to server")
			g.connected = true
		case EventDisconnectedFromServer:
			if g.connected {
				logger.FromContext(g.ctx).Info("Disconnected from server")
			} else {
				logger.FromContext(g.ctx).Errorf("Failed to connect to server")
			}
			g.connected = false
		default:
			logger.FromContext(g.ctx).Errorf("Unknown event: %d", event)
		}
	default:
		// Continue
	}

	const speed = 5
	if ebiten.IsKeyPressed(KeyLeft) {
		g.player.Pos.X -= speed
	} else if ebiten.IsKeyPressed(KeyRight) {
		g.player.Pos.X += speed
	}
	if ebiten.IsKeyPressed(KeyUp) {
		g.player.Pos.Y -= speed
	} else if ebiten.IsKeyPressed(KeyDown) {
		g.player.Pos.Y += speed
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.events <- EventStartServer
	}

	if g.connected {
		g.sendMessage(&socket.Request{
			Type: &socket.Request_PlayerMove{
				PlayerMove: &socket.Pos{
					X: g.player.Pos.X,
					Y: g.player.Pos.Y,
				},
			},
		})

		g.processMessages()
	}

	return nil
}

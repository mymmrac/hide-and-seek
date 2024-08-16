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

	const speed = 100.0
	dx, dy := 0.0, 0.0
	if ebiten.IsKeyPressed(KeyLeft) {
		dx -= speed
	} else if ebiten.IsKeyPressed(KeyRight) {
		dx += speed
	}
	if ebiten.IsKeyPressed(KeyUp) {
		dy -= speed
	} else if ebiten.IsKeyPressed(KeyDown) {
		dy += speed
	}

	g.player.Collider.Body().SetVelocity(dx, dy)

	// collisionPenetration := cp.Vector{}
	// g.player.Collider.Body().EachArbiter(func(arb *cp.Arbiter) {
	// 	penetration := arb.Normal().Mult(arb.ContactPointSet().Points[0].Distance)
	// 	if penetration.LengthSq() > collisionPenetration.LengthSq() {
	// 		collisionPenetration = penetration
	// 	}
	// })

	g.space.Step(1.0 / float64(ebiten.TPS()))

	pos := g.player.Collider.Body().Position()
	g.player.Pos.X = pos.X - g.player.Size.X/2
	g.player.Pos.Y = pos.Y - g.player.Size.Y/2

	g.camera.Position = g.player.Pos.Sub(g.camera.ViewportCenter())

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		if g.camera.Zoom > -2400 {
			g.camera.Zoom -= 1
		}
	} else if ebiten.IsKeyPressed(ebiten.KeyE) {
		if g.camera.Zoom < 2400 {
			g.camera.Zoom += 1
		}
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

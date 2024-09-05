package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"

	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
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

	g.ticks++

	if g.keybindings.IsActionJustPressed(ActionFullScreen) {
		ebiten.SetFullscreen(!ebiten.IsFullscreen())
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

	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		g.collisions = !g.collisions
	}

	const speed = 4.0
	move := space.Vec2F{}
	if g.keybindings.IsActionPressed(ActionWalkUp) {
		move.Y -= speed
		g.player.Dir.X = 0
		g.player.Dir.Y = -1
	} else if g.keybindings.IsActionPressed(ActionWalkDown) {
		move.Y += speed
		g.player.Dir.X = 0
		g.player.Dir.Y = 1
	}
	if g.keybindings.IsActionPressed(ActionWalkLeft) {
		move.X -= speed
		g.player.Dir.X = -1
		g.player.Dir.Y = 0
	} else if g.keybindings.IsActionPressed(ActionWalkRight) {
		move.X += speed
		g.player.Dir.X = 1
		g.player.Dir.Y = 0
	}

	g.player.Moving = !move.IsZero()

	coll := g.player.Collider
	if g.collisions {
		if collision := coll.Collide(move.OX()); collision != nil {
			move.X = collision.Resolve().X
		}

		if collision := coll.Collide(move.OY()); collision != nil {
			move.Y = collision.Resolve().Y
		}
	}

	coll.Move(move)
	g.player.UpdatePosition()

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
		pos := coll.Position()
		g.sendMessage(&socket.Request{
			Type: &socket.Request_PlayerState_{
				PlayerState: &socket.Request_PlayerState{
					Pos: &socket.Vec2F{
						X: pos.X,
						Y: pos.Y,
					},
					Dir: &socket.Vec2I{
						X: int32(g.player.Dir.X),
						Y: int32(g.player.Dir.Y),
					},
					Moving: g.player.Moving,
				},
			},
		})

		g.processMessages()
	}

	return nil
}

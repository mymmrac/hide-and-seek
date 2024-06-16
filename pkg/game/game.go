package game

import (
	"context"
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"

	"github.com/mymmrac/hide-and-seek/pkg/space"
)

type Game struct {
	ctx    context.Context
	cancel context.CancelFunc

	player Player
}

func NewGame(
	ctx context.Context,
	cancel context.CancelFunc,
) *Game {
	return &Game{
		ctx:    ctx,
		cancel: cancel,
		player: Player{
			Pos: space.Vec2F{
				X: 100,
				Y: 100,
			},
		},
	}
}

func (g *Game) Init() error {
	ebiten.SetWindowTitle("Hide & Seek")
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowClosingHandled(true)
	return nil
}

func (g *Game) Update() error {
	select {
	case <-g.ctx.Done():
		return ebiten.Termination
	default:
		// Continue
	}

	if ebiten.IsWindowBeingClosed() || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.cancel()
		return nil
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

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkgray)

	vector.DrawFilledRect(
		screen,
		float32(g.player.Pos.X),
		float32(g.player.Pos.Y),
		32,
		32,
		colornames.Blue,
		true,
	)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"FPS: %0.2f\nTPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()),
	)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 1080, 720
}

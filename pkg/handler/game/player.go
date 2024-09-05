package game

import (
	"image"
	"math/rand/v2"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/mymmrac/hide-and-seek/pkg/module/collider"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
)

type Player struct {
	Name     string
	Pos      space.Vec2F
	Dir      space.Vec2I
	Moving   bool
	Collider *collider.Object
}

func NewPlayer(cw *collider.World) *Player {
	return &Player{
		Name:     "test" + strconv.FormatUint(rand.Uint64N(9000)+1000, 10),
		Pos:      space.Vec2F{},
		Dir:      space.Vec2I{X: 0, Y: 1},
		Collider: cw.NewObject(playerColliderOffset, space.Vec2F{X: 24, Y: 24}),
	}
}

func (p *Player) UpdatePosition() {
	pos := p.Collider.Position().Sub(playerColliderOffset)
	p.Pos = pos
}

var playerColliderOffset = space.Vec2F{X: 4, Y: 8}

func (p *Player) Draw(screen, spriteSheet *ebiten.Image, ticks int) DrawCall {
	var dx, dy int

	const frameCount = 6
	switch p.Dir {
	case space.Vec2I{X: 1, Y: 0}:
		dx = 0 * frameCount
	case space.Vec2I{X: -1, Y: 0}:
		dx = 2 * frameCount
	case space.Vec2I{X: 0, Y: 1}:
		dx = 3 * frameCount
	case space.Vec2I{X: 0, Y: -1}:
		dx = 1 * frameCount
	default:
		panic("unreachable")
	}

	var animationSpeed int
	if p.Moving {
		dy = 2
		animationSpeed = 10
	} else {
		dy = 1
		animationSpeed = 16
	}
	dx += (ticks / animationSpeed) % frameCount

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Pos.X, p.Pos.Y-32)
	playerImg := spriteSheet.SubImage(image.Rect(32*dx, 64*dy, 32*(dx+1), 64*(dy+1))).(*ebiten.Image)

	return DrawCall{
		Y: p.Pos.Y - 32,
		F: func() {
			screen.DrawImage(playerImg, op)
			ebitenutil.DebugPrintAt(screen, p.Name, int(p.Pos.X), int(p.Pos.Y)-32)
		},
	}
}

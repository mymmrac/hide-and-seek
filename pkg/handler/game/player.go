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

func (p *Player) Draw(screen, spriteSheet *ebiten.Image) {
	var playerSprite image.Rectangle
	switch p.Dir {
	case space.Vec2I{X: 1, Y: 0}:
		playerSprite = image.Rect(32*0, 0, 32*1, 64)
	case space.Vec2I{X: -1, Y: 0}:
		playerSprite = image.Rect(32*2, 0, 32*3, 64)
	case space.Vec2I{X: 0, Y: 1}:
		playerSprite = image.Rect(32*3, 0, 32*4, 64)
	case space.Vec2I{X: 0, Y: -1}:
		playerSprite = image.Rect(32*1, 0, 32*2, 64)
	default:
		panic("unreachable")
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.Pos.X, p.Pos.Y-32)
	playerImg := spriteSheet.SubImage(playerSprite).(*ebiten.Image)
	screen.DrawImage(playerImg, op)

	ebitenutil.DebugPrintAt(screen, p.Name, int(p.Pos.X), int(p.Pos.Y)-32)
}

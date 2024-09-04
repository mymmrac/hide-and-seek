package game

import (
	"math/rand/v2"
	"strconv"

	"github.com/mymmrac/hide-and-seek/pkg/module/collider"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
)

type Player struct {
	Name     string
	Pos      space.Vec2F
	Size     space.Vec2F
	Dir      space.Vec2I
	Collider *collider.Object
}

func NewPlayer(cw *collider.World) *Player {
	return &Player{
		Name:     "test" + strconv.FormatUint(rand.Uint64N(9000)+1000, 10),
		Pos:      space.Vec2F{},
		Size:     space.Vec2F{X: 32, Y: 32},
		Dir:      space.Vec2I{X: 0, Y: 1},
		Collider: cw.NewObject(playerColliderOffset, space.Vec2F{X: 24, Y: 24}),
	}
}

func (p *Player) UpdatePosition() {
	pos := p.Collider.Position().Sub(playerColliderOffset)
	p.Pos = pos
}

var playerColliderOffset = space.Vec2F{X: 4, Y: 8}

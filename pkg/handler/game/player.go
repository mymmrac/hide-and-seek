package game

import (
	"github.com/solarlune/resolv"

	"github.com/mymmrac/hide-and-seek/pkg/module/space"
)

type Player struct {
	Name     string
	Pos      space.Vec2F
	Collider *resolv.Object
}

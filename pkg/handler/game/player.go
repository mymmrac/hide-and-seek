package game

import (
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

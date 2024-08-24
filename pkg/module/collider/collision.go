package collider

import "github.com/mymmrac/hide-and-seek/pkg/module/space"

type Collision struct {
	Resolutions []CollisionResolution
}

func (c Collision) Resolve() space.Vec2F {
	return c.Resolutions[0].Delta
}

type CollisionResolution struct {
	Object *Object
	Delta  space.Vec2F
}

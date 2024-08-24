package collider

import (
	"unique"

	"github.com/mymmrac/hide-and-seek/pkg/module/space"
	"github.com/mymmrac/hide-and-seek/pkg/module/stream"
)

type Object struct {
	world *World

	pos  space.Vec2F
	size space.Vec2F

	tags []unique.Handle[string]
}

func NewObject(pos space.Vec2F, size space.Vec2F, tags ...string) *Object {
	if size.X <= 0 || size.Y <= 0 {
		panic("size must be positive")
	}
	return &Object{
		world: nil,
		pos:   pos,
		size:  size,
		tags:  stream.Collect(stream.Map(stream.Stream(tags), unique.Make)),
	}
}

func (o *Object) SetPosition(pos space.Vec2F) {
	o.pos = pos
}

func (o *Object) Position() space.Vec2F {
	return o.pos
}

func (o *Object) Size() space.Vec2F {
	return o.size
}

// Center returns the center of the object
func (o *Object) Center() space.Vec2F {
	return o.pos.Add(o.size.Scale(0.5))
}

// Collide returns the collision if the object collided with the other objects
func (o *Object) Collide(move space.Vec2F, tags ...string) *Collision {
	return o.world.Collide(o, move, tags...)
}

// CollideWith returns a delta vector if the object collided with the other object
func (o *Object) CollideWith(other *Object, move space.Vec2F) *space.Vec2F {
	if collision := o.StaticCollisionWith(other); collision != nil {
		return collision
	}

	if move.Len2() == 0 {
		return nil
	}

	mo := &Object{pos: o.pos.Add(move), size: o.size}
	delta := mo.StaticCollisionWith(other)
	if delta == nil {
		return nil
	}

	var dx float64
	if o.pos.X > other.pos.X+other.size.X {
		dx = o.pos.X - (other.pos.X + other.size.X)
	} else if other.pos.X > o.pos.X+o.size.X {
		dx = other.pos.X - (o.pos.X + o.size.X)
	}

	var dy float64
	if o.pos.Y > other.pos.Y+other.size.Y {
		dy = o.pos.Y - (other.pos.Y + other.size.Y)
	} else if other.pos.Y > o.pos.Y+o.size.Y {
		dy = other.pos.Y - (o.pos.Y + o.size.Y)
	}

	return &space.Vec2F{X: dx * sign(move.X), Y: dy * sign(move.Y)}
}

func sign(x float64) float64 {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func (o *Object) StaticCollisionWith(other *Object) *space.Vec2F {
	if !o.Intercepts(other) {
		return nil
	}

	dir := o.Center().Sub(other.Center()).Norm()

	minX := min((o.pos.X+o.size.X)-other.pos.X, (other.pos.X+other.size.X)-o.pos.X)
	minY := min((o.pos.Y+o.size.Y)-other.pos.Y, (other.pos.Y+other.size.Y)-o.pos.Y)

	return &space.Vec2F{
		X: dir.X * min(minX, minY),
		Y: dir.Y * min(minX, minY),
	}
}

// ContainsPoint returns true if the object contains the point
func (o *Object) ContainsPoint(point space.Vec2F) bool {
	return o.pos.X < point.X && point.X < o.pos.X+o.size.X &&
		o.pos.Y < point.Y && point.Y < o.pos.Y+o.size.Y
}

// Intercepts returns true if the object intercepts the other object
func (o *Object) Intercepts(other *Object) bool {
	return o.pos.X < other.pos.X+other.size.X && other.pos.X < o.pos.X+o.size.X &&
		o.pos.Y < other.pos.Y+other.size.Y && other.pos.Y < o.pos.Y+o.size.Y
}

package collider

import (
	"cmp"
	"slices"
	"unique"

	"github.com/mymmrac/hide-and-seek/pkg/module/collection"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
	"github.com/mymmrac/hide-and-seek/pkg/module/stream"
)

type World struct {
	objects []*Object
}

func NewWorld() *World {
	return &World{}
}

func (w *World) Add(objects ...*Object) {
	for _, object := range objects {
		object.world = w
	}
	w.objects = append(w.objects, objects...)
}

func (w *World) Remove(objects ...*Object) {
	for _, object := range objects {
		object.world = nil
	}
	w.objects = slices.DeleteFunc(w.objects, func(o *Object) bool {
		return slices.Contains(objects, o)
	})
}

func (w *World) NewObject(pos space.Vec2F, size space.Vec2F, tags ...string) *Object {
	obj := NewObject(pos, size, tags...)
	obj.world = w
	w.objects = append(w.objects, obj)
	return obj
}

func (w *World) Objects() []*Object {
	return w.objects
}

func (w *World) Collide(obj *Object, move space.Vec2F, tags ...string) *Collision {
	objects := stream.Filter(stream.Stream(w.objects), func(o *Object) bool {
		return o != obj && collection.ContainsAll(o.tags, stream.Collect(stream.Map(stream.Stream(tags), unique.Make))...)
	})

	var collision *Collision
	for object := range objects {
		if delta := obj.CollideWith(object, move); delta != nil {
			if collision == nil {
				collision = &Collision{
					Resolutions: []CollisionResolution{{
						Object: object,
						Delta:  *delta,
					}},
				}
			} else {
				collision.Resolutions = append(collision.Resolutions, CollisionResolution{
					Object: object,
					Delta:  *delta,
				})
			}
		}
	}

	if collision != nil {
		slices.SortFunc(collision.Resolutions, func(a, b CollisionResolution) int {
			return cmp.Compare(b.Delta.Len2(), a.Delta.Len2())
		})
	}

	return collision
}

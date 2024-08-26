package world

import "github.com/mymmrac/hide-and-seek/pkg/module/space"

type Defs struct {
	Tilesets map[int]TilesetDef
	Entities map[int]EntityDef
}

type TilesetDef struct {
	Path  string
	Tiles map[int]TileDef
}

type TileDef struct {
	Pos  space.Vec2I
	Size space.Vec2I
}

type EntityDef struct {
	TilesetID int
	TileID    int
	Size      space.Vec2I
}

type World struct {
	Levels []Level
	Spawn  space.Vec2I
}

type Level struct {
	Pos      space.Vec2I
	Tiles    []Tile
	WallSize space.Vec2I
	Walls    []space.Vec2I
	Entities []Entity
}

type Tile struct {
	TilesetID int
	TileID    int
	Pos       space.Vec2I
}

type Entity struct {
	EntityID int
	Pos      space.Vec2I
}

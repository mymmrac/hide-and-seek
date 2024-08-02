package world

import "github.com/mymmrac/hide-and-seek/pkg/module/space"

type Defs struct {
	Tilesets map[int]Tileset
}

type Tileset struct {
	Path     string
	TileSize space.Vec2I
	Tiles    map[int]space.Vec2I
}

type World struct {
	Levels []Level
	Spawn  space.Vec2I
}

type Level struct {
	Pos   space.Vec2I
	Tiles []Tile
}

type Tile struct {
	Pos       space.Vec2I
	TilesetID int
	TileID    int
}

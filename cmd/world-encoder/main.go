package main

import (
	"cmp"
	"context"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
	"github.com/mymmrac/hide-and-seek/pkg/module/world"
)

func main() {
	ctx := context.Background()
	log := logger.FromContext(ctx)

	err := encodeWorlds(log, "./assets/world/worlds.ldtk", "./assets/world/")
	if err != nil {
		log.Errorf("Error encoding world: %s", err)
		os.Exit(1)
	}
}

func encodeWorlds(log *logger.Logger, ldtkFilePath, outputDirPath string) error {
	ldtkFile, err := os.Open(ldtkFilePath)
	if err != nil {
		return fmt.Errorf("open LDtk file: %w", err)
	}

	var ldtk LDtkFile
	err = json.NewDecoder(ldtkFile).Decode(&ldtk)
	if err != nil {
		return fmt.Errorf("decode LDtk file: %w", err)
	}

	defs := world.Defs{
		Tilesets: make(map[int]world.TilesetDef, len(ldtk.Defs.Tilesets)),
		Entities: make(map[int]world.EntityDef),
	}

	tileSizes := make(map[int]space.Vec2I)
	for _, tileset := range ldtk.Defs.Tilesets {
		defs.Tilesets[tileset.UID] = world.TilesetDef{
			Path:  tileset.RelPath,
			Tiles: make(map[int]world.TileDef),
		}
		tileSizes[tileset.UID] = space.Vec2I{
			X: tileset.TileGridSize,
			Y: tileset.TileGridSize,
		}
	}

	lastTileID := 1
	findOrAddTile := func(tilesetID int, tileDef world.TileDef) (int, error) {
		tileset, ok := defs.Tilesets[tilesetID]
		if !ok {
			return 0, fmt.Errorf("tileset %d not found", tilesetID)
		}

		var tileID int
		for otherTileID, otherTileDef := range tileset.Tiles {
			if otherTileDef == tileDef {
				tileID = otherTileID
				break
			}
		}
		if tileID == 0 {
			tileID = lastTileID
			tileset.Tiles[lastTileID] = tileDef
			lastTileID++
		}

		return tileID, nil
	}

	for _, entity := range ldtk.Defs.Entities {
		var tileID int
		tileID, err = findOrAddTile(entity.TilesetID, world.TileDef{
			Pos: space.Vec2I{
				X: entity.TileRect.X,
				Y: entity.TileRect.Y,
			},
			Size: space.Vec2I{
				X: entity.TileRect.W,
				Y: entity.TileRect.H,
			},
		})
		if err != nil {
			return err
		}

		defs.Entities[entity.UID] = world.EntityDef{
			TilesetID: entity.TilesetID,
			TileID:    tileID,
			Size: space.Vec2I{
				X: entity.Width,
				Y: entity.Height,
			},
		}
	}

	for _, w := range ldtk.Worlds {
		wd := world.World{
			Levels: make([]world.Level, len(w.Levels)),
			Spawn:  space.Vec2I{},
		}

		for i, lvl := range w.Levels {
			lv := world.Level{
				Pos: space.Vec2I{
					X: lvl.WorldX,
					Y: lvl.WorldY,
				},
				Tiles:    nil,
				WallSize: space.Vec2I{},
				Walls:    nil,
				Entities: nil,
			}
			for _, layer := range lvl.LayerInstances {
				switch layer.Identifier {
				case "walls_and_floor":
					for _, tile := range layer.AutoLayerTiles {
						tileDef := world.TileDef{
							Pos: space.Vec2I{
								X: tile.Src[0],
								Y: tile.Src[1],
							},
							Size: tileSizes[layer.TilesetDefUID],
						}

						var tileID int
						tileID, err = findOrAddTile(layer.TilesetDefUID, tileDef)
						if err != nil {
							return err
						}

						lv.Tiles = append(lv.Tiles, world.Tile{
							TilesetID: layer.TilesetDefUID,
							TileID:    tileID,
							Pos: space.Vec2I{
								X: tile.Px[0],
								Y: tile.Px[1],
							},
						})
					}
				case "layout":
					k := slices.IndexFunc(ldtk.Defs.Layers, func(layer Layers) bool {
						return layer.Identifier == "layout"
					})
					if k < 0 {
						return fmt.Errorf("layout layer not found")
					}
					layoutLayer := ldtk.Defs.Layers[k]

					wallValue := -1
					bottomWallValue := -1
					for _, value := range layoutLayer.IntGridValues {
						if value.Identifier == "wall" {
							wallValue = value.Value
						}
						if value.Identifier == "bottom_wall" {
							bottomWallValue = value.Value
						}
					}
					if wallValue < 0 {
						return fmt.Errorf("wall values not found")
					}
					if bottomWallValue < 0 {
						return fmt.Errorf("bottom wall values not found")
					}

					lv.WallSize = space.Vec2I{
						X: layer.GridSize,
						Y: layer.GridSize,
					}

					for j, value := range layer.IntGridCsv {
						switch value {
						case wallValue:
							lv.Walls = append(lv.Walls, space.Vec2I{
								X: j % layer.CWid,
								Y: j / layer.CWid,
							})
						case bottomWallValue:
							lv.Walls = append(lv.Walls, space.Vec2I{
								X: j % layer.CWid,
								Y: j / layer.CWid,
							})
						}
					}
				case "furniture_entities":
					for _, entity := range layer.EntityInstances {
						lv.Entities = append(lv.Entities, world.Entity{
							EntityID: entity.DefUID,
							Pos: space.Vec2I{
								X: entity.WorldX,
								Y: entity.WorldY,
							},
						})
					}
				case "entities":
					for _, entity := range layer.EntityInstances {
						if entity.Identifier == "spawn" {
							wd.Spawn = space.Vec2I{
								X: entity.WorldX,
								Y: entity.WorldY,
							}
						}
					}
				default:
					continue
				}
			}

			slices.SortFunc(lv.Tiles, func(a, b world.Tile) int {
				if a.TilesetID == b.TilesetID {
					return cmp.Compare(a.TileID, b.TileID)
				}
				return cmp.Compare(a.TilesetID, b.TilesetID)
			})

			slices.SortFunc(lv.Entities, func(a, b world.Entity) int {
				return cmp.Compare(a.EntityID, b.EntityID)
			})

			wd.Levels[i] = lv
		}

		if err = encode(defs, filepath.Join(outputDirPath, "defs.bin")); err != nil {
			return fmt.Errorf("defs: %w", err)
		}

		if err = encode(wd, filepath.Join(outputDirPath, fmt.Sprintf("world_%s.bin", w.Identifier))); err != nil {
			return fmt.Errorf("world: %w", err)
		}
	}

	return nil
}

func encode(v any, path string) error {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("create file: %w", err)
	}

	err = gob.NewEncoder(file).Encode(v)
	if err != nil {
		return fmt.Errorf("encode: %w", err)
	}

	if err = file.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}

	return nil
}

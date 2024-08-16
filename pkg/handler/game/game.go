package game

import (
	"context"
	"encoding/gob"
	"fmt"
	"image"
	"io/fs"
	"math/rand/v2"
	"strconv"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp/v2"

	"github.com/mymmrac/hide-and-seek/assets"
	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/camera"
	"github.com/mymmrac/hide-and-seek/pkg/module/collection"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
	"github.com/mymmrac/hide-and-seek/pkg/module/world"
)

type Game struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	httpClient *fiber.Client

	events chan EventType

	connected    bool
	connectionID uint64
	requests     chan *socket.Request
	responses    chan *socket.Response

	camera   *camera.Camera
	worldImg *ebiten.Image

	defs  world.Defs
	world world.World

	tilesets map[int]*ebiten.Image

	players *collection.SyncMap[uint64, *Player]

	info *socket.Response_Info

	player Player

	space *cp.Space
}

func NewGame(
	ctx context.Context,
	cancel context.CancelFunc,
) *Game {
	sp := cp.NewSpace()

	playerBody := sp.AddBody(cp.NewBody(1, cp.INFINITY))

	playerShape := sp.AddShape(cp.NewBox(playerBody, 32, 32, 1))
	playerShape.SetElasticity(1)
	playerShape.SetFriction(1)

	return &Game{
		ctx:          ctx,
		cancel:       cancel,
		wg:           sync.WaitGroup{},
		httpClient:   &fiber.Client{},
		events:       make(chan EventType, 32),
		connected:    false,
		connectionID: rand.Uint64(),
		requests:     nil,
		responses:    nil,
		camera:       &camera.Camera{Viewport: space.Vec2F{X: 1080, Y: 720}, Zoom: 100},
		worldImg:     ebiten.NewImage(2048, 2048),
		defs:         world.Defs{},
		world:        world.World{},
		tilesets:     make(map[int]*ebiten.Image),
		players:      collection.NewSyncMap[uint64, *Player](),
		info:         nil,
		player: Player{
			Name:     "test" + strconv.FormatUint(rand.Uint64N(9000)+1000, 10),
			Pos:      space.Vec2F{},
			Size:     space.Vec2F{X: 32, Y: 32},
			Collider: playerShape,
		},
		space: sp,
	}
}

func (g *Game) Init() error {
	g.wg.Add(1) // WG: Game loop

	ebiten.SetWindowTitle("Hide & Seek")
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetVsyncEnabled(false)

	defsFile, err := assets.FS.Open("world/defs.bin")
	if err != nil {
		return fmt.Errorf("load defs file: %w", err)
	}

	if err = gob.NewDecoder(defsFile).Decode(&g.defs); err != nil {
		return fmt.Errorf("decode defs: %w", err)
	}

	for id, tileset := range g.defs.Tilesets {
		if tileset.Path == "" {
			logger.FromContext(g.ctx).Debugf("Skipping tileset %d, empty path", id)
			continue
		}

		var tilesetImageFile fs.File
		tilesetImageFile, err = assets.FS.Open(strings.TrimPrefix(tileset.Path, "../"))
		if err != nil {
			return fmt.Errorf("open %d %q tileset: %w", id, tileset.Path, err)
		}

		var tilesetImage image.Image
		tilesetImage, _, err = image.Decode(tilesetImageFile)
		if err != nil {
			return fmt.Errorf("decode %d tileset image: %w", id, err)
		}

		g.tilesets[id] = ebiten.NewImageFromImage(tilesetImage)
	}

	worldFile, err := assets.FS.Open("world/world_office_0.bin")
	if err != nil {
		return fmt.Errorf("load world file: %w", err)
	}

	if err = gob.NewDecoder(worldFile).Decode(&g.world); err != nil {
		return fmt.Errorf("decode world: %w", err)
	}

	g.player.Pos = g.world.Spawn.ToF()
	g.player.Collider.Body().SetPosition(cp.Vector{
		X: g.player.Pos.X + g.player.Size.X/2,
		Y: g.player.Pos.Y + g.player.Size.Y/2,
	})

	for _, level := range g.world.Levels {
		for _, wall := range level.Walls {
			body := g.space.AddBody(cp.NewStaticBody())
			body.SetPosition(cp.Vector{
				X: float64(wall.X*level.WallSize.X + level.Pos.X + level.WallSize.X/2),
				Y: float64(wall.Y*level.WallSize.Y + level.Pos.Y + level.WallSize.Y/2),
			})

			shape := g.space.AddShape(cp.NewBox(body, float64(level.WallSize.X), float64(level.WallSize.Y), 1))
			shape.SetElasticity(0)
			shape.SetFriction(0)
		}
	}

	return nil
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func (g *Game) Shutdown() {
	g.cancel()
	g.wg.Wait()
}

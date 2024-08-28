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

	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mymmrac/hide-and-seek/assets"
	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/camera"
	"github.com/mymmrac/hide-and-seek/pkg/module/chttp"
	"github.com/mymmrac/hide-and-seek/pkg/module/collection"
	"github.com/mymmrac/hide-and-seek/pkg/module/collider"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
	"github.com/mymmrac/hide-and-seek/pkg/module/world"
)

const (
	defaultScreenWidth  = 1280
	defaultScreenHeight = 720
)

type Game struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	httpClient chttp.Client

	events chan EventType

	connected    bool
	connectionID uint64
	requests     chan *socket.Request
	responses    chan *socket.Response

	keybindings KeyBindings

	camera   *camera.Camera
	worldImg *ebiten.Image

	defs  world.Defs
	world world.World

	tilesets map[int]*ebiten.Image

	players *collection.SyncMap[uint64, *Player]

	info *socket.Response_Info

	player Player

	collisions bool
	cw         *collider.World
}

func NewGame(
	ctx context.Context,
	cancel context.CancelFunc,
) *Game {
	cw := collider.NewWorld()
	return &Game{
		ctx:          ctx,
		cancel:       cancel,
		wg:           sync.WaitGroup{},
		httpClient:   chttp.DefaultClient,
		events:       make(chan EventType, 32),
		connected:    false,
		connectionID: rand.Uint64(),
		requests:     nil,
		responses:    nil,
		keybindings:  DefaultKeyBindings.Clone(),
		camera: &camera.Camera{
			Viewport: space.Vec2F{X: defaultScreenWidth, Y: defaultScreenHeight},
			Zoom:     100,
		},
		worldImg: ebiten.NewImage(2048, 2048),
		defs:     world.Defs{},
		world:    world.World{},
		tilesets: make(map[int]*ebiten.Image),
		players:  collection.NewSyncMap[uint64, *Player](),
		info:     nil,
		player: Player{
			Name:     "test" + strconv.FormatUint(rand.Uint64N(9000)+1000, 10),
			Pos:      space.Vec2F{},
			Size:     space.Vec2F{X: 32, Y: 32},
			Collider: cw.NewObject(space.Vec2F{}, space.Vec2F{X: 32, Y: 32}),
		},
		collisions: true,
		cw:         cw,
	}
}

func (g *Game) Init() error {
	g.wg.Add(1) // WG: Game loop

	ebiten.SetWindowTitle("Hide & Seek")
	ebiten.SetWindowSize(defaultScreenWidth, defaultScreenHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowClosingHandled(true)
	ebiten.SetVsyncEnabled(false)
	ebiten.SetScreenClearedEveryFrame(false)

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
		log.Debugf("Loaded tileset %d %q", id, tileset.Path)
	}

	worldFile, err := assets.FS.Open("world/world_office_0.bin")
	if err != nil {
		return fmt.Errorf("load world file: %w", err)
	}

	if err = gob.NewDecoder(worldFile).Decode(&g.world); err != nil {
		return fmt.Errorf("decode world: %w", err)
	}

	g.player.Pos = g.world.Spawn.ToF()
	g.player.Collider.SetPosition(g.player.Pos)

	for _, level := range g.world.Levels {
		for _, coll := range level.Colliders {
			g.cw.NewObject(coll.Pos.ToF(), coll.Size.ToF(), coll.Tags...)
		}
	}

	return nil
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return defaultScreenWidth, defaultScreenHeight
}

func (g *Game) Shutdown() {
	g.cancel()
	g.wg.Wait()
}

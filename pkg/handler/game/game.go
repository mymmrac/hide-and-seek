package game

import (
	"context"
	"encoding/gob"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"sync"

	"github.com/charmbracelet/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"

	"github.com/mymmrac/hide-and-seek/assets"
	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/camera"
	"github.com/mymmrac/hide-and-seek/pkg/module/chttp"
	"github.com/mymmrac/hide-and-seek/pkg/module/collection"
	"github.com/mymmrac/hide-and-seek/pkg/module/collider"
	"github.com/mymmrac/hide-and-seek/pkg/module/loader"
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

	httpClient    chttp.Client
	serverAddress string

	events chan EventType

	connected    bool
	connectionID uint64
	requests     chan *socket.Request
	responses    chan *socket.Response

	keybindings KeyBindings

	camera   *camera.Camera
	worldImg *ebiten.Image
	ticks    int

	defs  world.Defs
	world world.World

	tilesets map[int]*ebiten.Image

	players *collection.SyncMap[uint64, *Player]

	info *socket.Response_Info

	player            *Player
	playerSpriteSheet *ebiten.Image

	collisions bool
	cw         *collider.World

	audioCtx        *audio.Context
	backgroundMusic *audio.Player
}

func NewGame(
	ctx context.Context,
	cancel context.CancelFunc,
) *Game {
	cw := collider.NewWorld()
	return &Game{
		ctx:           ctx,
		cancel:        cancel,
		wg:            sync.WaitGroup{},
		httpClient:    chttp.DefaultClient,
		serverAddress: "localhost:4242",
		events:        make(chan EventType, 32),
		connected:     false,
		connectionID:  rand.Uint64(),
		requests:      nil,
		responses:     nil,
		keybindings:   DefaultKeyBindings.Clone(),
		camera: &camera.Camera{
			Viewport: space.Vec2F{X: defaultScreenWidth, Y: defaultScreenHeight},
			Zoom:     100,
		},
		worldImg:          ebiten.NewImage(2048, 2048),
		ticks:             0,
		defs:              world.Defs{},
		world:             world.World{},
		tilesets:          make(map[int]*ebiten.Image),
		players:           collection.NewSyncMap[uint64, *Player](),
		info:              nil,
		player:            NewPlayer(cw),
		playerSpriteSheet: nil,
		collisions:        true,
		cw:                cw,
		audioCtx:          audio.NewContext(44100),
		backgroundMusic:   nil,
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

	if len(os.Args) == 2 {
		g.serverAddress = os.Args[1]
	}

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

		var tilesetImage *ebiten.Image
		tilesetImage, err = loader.Image(strings.TrimPrefix(tileset.Path, "../"))
		if err != nil {
			return fmt.Errorf("load %d tileset image: %w", id, err)
		}

		g.tilesets[id] = tilesetImage
		log.Debugf("Loaded tileset %d %q", id, tileset.Path)
	}

	worldFile, err := assets.FS.Open("world/world_office_0.bin")
	if err != nil {
		return fmt.Errorf("load world file: %w", err)
	}

	if err = gob.NewDecoder(worldFile).Decode(&g.world); err != nil {
		return fmt.Errorf("decode world: %w", err)
	}

	g.playerSpriteSheet, err = loader.Image("images/Premade_Character_32x32_19.png")
	if err != nil {
		return fmt.Errorf("load player sprite sheet: %w", err)
	}

	g.player.Collider.SetPosition(g.world.Spawn.ToF().Add(playerColliderOffset))
	g.player.UpdatePosition()

	for _, level := range g.world.Levels {
		for _, coll := range level.Colliders {
			g.cw.NewObject(coll.Pos.ToF(), coll.Size.ToF(), coll.Tags...)
		}

		for _, entity := range level.Entities {
			entityDef := g.defs.Entities[entity.EntityID]
			for _, coll := range entityDef.Colliders {
				g.cw.NewObject(
					entity.Pos.ToF().Add(coll.Pos.ToF()),
					coll.Size.ToF(),
					coll.Tags...,
				)
			}
		}
	}

	g.backgroundMusic, err = loader.Audio(g.audioCtx, "music/Towns/Bustling Streets.ogg", true)
	if err != nil {
		return fmt.Errorf("load background music player: %w", err)
	}

	g.backgroundMusic.SetVolume(0.5)
	g.backgroundMusic.Play()

	return nil
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return defaultScreenWidth, defaultScreenHeight
}

func (g *Game) Shutdown() {
	g.cancel()
	g.wg.Wait()
}

package game

import (
	"context"
	"math/rand/v2"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/hajimehoshi/ebiten/v2"

	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/collection"
	"github.com/mymmrac/hide-and-seek/pkg/module/space"
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

	players *collection.SyncMap[uint64, *Player]

	info *socket.Response_Info

	player Player
}

func NewGame(
	ctx context.Context,
	cancel context.CancelFunc,
) *Game {
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
		players:      collection.NewSyncMap[uint64, *Player](),
		info:         nil,
		player: Player{
			Name: "test" + strconv.FormatUint(rand.Uint64N(9000)+1000, 10),
			Pos: space.Vec2F{
				X: 100,
				Y: 100,
			},
		},
	}
}

func (g *Game) Init() error {
	g.wg.Add(1) // WG: Game loop

	ebiten.SetWindowTitle("Hide & Seek")
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowClosingHandled(true)

	return nil
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func (g *Game) Shutdown() {
	g.cancel()
	g.wg.Wait()
}

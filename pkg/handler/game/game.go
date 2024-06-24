package game

import (
	"context"
	"fmt"
	"math/rand/v2"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"

	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
	"github.com/mymmrac/hide-and-seek/pkg/module/collection"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
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

	g.events <- EventStartServer

	return nil
}

func (g *Game) Update() error {
	select {
	case <-g.ctx.Done():
		g.wg.Done() // WG: Game loop
		return ebiten.Termination
	default:
		// Continue
	}

	if ebiten.IsWindowBeingClosed() || inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.cancel()
		return nil
	}

	select {
	case event := <-g.events:
		switch event {
		case EventStartServer:
			go g.connectToServer()
		case EventConnectedToServer:
			logger.FromContext(g.ctx).Info("Connected to server")
			g.connected = true
		case EventDisconnectedFromServer:
			if g.connected {
				logger.FromContext(g.ctx).Info("Disconnected from server")
			} else {
				logger.FromContext(g.ctx).Errorf("Failed to connect to server")
			}
			g.connected = false
		default:
			logger.FromContext(g.ctx).Errorf("Unknown event: %d", event)
		}
	default:
		// Continue
	}

	if g.connected {
		g.sendMessage(&socket.Request{
			Type: &socket.Request_PlayerMove{
				PlayerMove: &socket.Pos{
					X: g.player.Pos.X,
					Y: g.player.Pos.Y,
				},
			},
		})

		g.processMessages()
	}

	if g.info == nil {
		return nil
	}

	const speed = 5
	if ebiten.IsKeyPressed(KeyLeft) {
		g.player.Pos.X -= speed
	} else if ebiten.IsKeyPressed(KeyRight) {
		g.player.Pos.X += speed
	}
	if ebiten.IsKeyPressed(KeyUp) {
		g.player.Pos.Y -= speed
	} else if ebiten.IsKeyPressed(KeyDown) {
		g.player.Pos.Y += speed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkgray)

	if g.info == nil {
		return
	}

	vector.DrawFilledRect(
		screen,
		float32(g.player.Pos.X),
		float32(g.player.Pos.Y),
		32,
		32,
		colornames.Blue,
		true,
	)
	ebitenutil.DebugPrintAt(screen, g.player.Name, int(g.player.Pos.X), int(g.player.Pos.Y))

	g.players.ForEach(func(_ uint64, player *Player) bool {
		vector.DrawFilledRect(
			screen,
			float32(player.Pos.X),
			float32(player.Pos.Y),
			32,
			32,
			colornames.Green,
			true,
		)
		ebitenutil.DebugPrintAt(screen, player.Name, int(player.Pos.X), int(player.Pos.Y))
		return true
	})

	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Connected: %t", g.connected), 10, 100)

	ebitenutil.DebugPrint(screen, fmt.Sprintf(
		"FPS: %0.2f\nTPS: %0.2f", ebiten.ActualFPS(), ebiten.ActualTPS()),
	)
}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 1080, 720
}

func (g *Game) Shutdown() {
	g.cancel()
	g.wg.Wait()
}

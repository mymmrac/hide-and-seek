package game

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"

	"github.com/mymmrac/hide-and-seek/pkg/api"
	"github.com/mymmrac/hide-and-seek/pkg/logger"
	"github.com/mymmrac/hide-and-seek/pkg/space"
)

type Msg struct {
	Data []byte
}

type Game struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup

	events chan EventType

	connected    bool
	connectionID uint64
	connWrite    chan *api.Msg
	connRead     chan *api.Msg

	playerLock sync.RWMutex
	players    map[uint64]space.Vec2F

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
		events:       make(chan EventType, 32),
		connected:    false,
		connectionID: rand.Uint64(),
		connWrite:    nil,
		connRead:     nil,
		playerLock:   sync.RWMutex{},
		players:      make(map[uint64]space.Vec2F),
		player: Player{
			Pos: space.Vec2F{
				X: 100,
				Y: 100,
			},
		},
	}
}

func (g *Game) Init() error {
	g.wg.Add(1)

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
		g.wg.Done()
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
			go g.ConnectToServer()
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

	if g.connected {
		select {
		case g.connWrite <- &api.Msg{
			From: g.connectionID,
			Pos:  g.player.Pos,
		}:
			// Continue
		default:
			logger.FromContext(g.ctx).Errorf("Write buffer is full")
		}

		g.playerLock.Lock()
	msgLoop:
		for {
			select {
			case msg := <-g.connRead:
				g.players[msg.From] = msg.Pos
			default:
				break msgLoop
			}
		}
		g.playerLock.Unlock()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Darkgray)

	vector.DrawFilledRect(
		screen,
		float32(g.player.Pos.X),
		float32(g.player.Pos.Y),
		32,
		32,
		colornames.Blue,
		true,
	)

	g.playerLock.RLock()
	for _, playerPos := range g.players {
		vector.DrawFilledRect(
			screen,
			float32(playerPos.X),
			float32(playerPos.Y),
			32,
			32,
			colornames.Green,
			true,
		)
	}
	g.playerLock.RUnlock()

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

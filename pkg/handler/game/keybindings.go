package game

import (
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Action uint

const (
	_ Action = iota
	ActionWalkUp
	ActionWalkDown
	ActionWalkLeft
	ActionWalkRight
)

type KeyBindings map[Action][]ebiten.Key

var DefaultKeyBindings = KeyBindings{
	ActionWalkUp:    []ebiten.Key{ebiten.KeyW, ebiten.KeyUp},
	ActionWalkDown:  []ebiten.Key{ebiten.KeyS, ebiten.KeyDown},
	ActionWalkLeft:  []ebiten.Key{ebiten.KeyA, ebiten.KeyLeft},
	ActionWalkRight: []ebiten.Key{ebiten.KeyD, ebiten.KeyRight},
}

func (k KeyBindings) Clone() KeyBindings {
	nk := make(KeyBindings, len(k))
	for action, keys := range k {
		nk[action] = slices.Clone(keys)
	}
	return nk
}

func (k KeyBindings) IsActionPressed(action Action) bool {
	keys, ok := k[action]
	if !ok {
		return false
	}

	for _, key := range keys {
		if ebiten.IsKeyPressed(key) {
			return true
		}
	}

	return false
}

func (k KeyBindings) IsActionJustPressed(action Action) bool {
	keys, ok := k[action]
	if !ok {
		return false
	}

	for _, key := range keys {
		if inpututil.IsKeyJustPressed(key) {
			return true
		}
	}

	return false
}

package game

type EventType uint

const (
	_ EventType = iota

	EventStartServer
	EventConnectedToServer
	EventDisconnectedFromServer
)

package server

import (
	"github.com/mymmrac/hide-and-seek/pkg/api/socket"
)

type Client struct {
	PlayerID     uint64
	PlayerName   string
	ConnectionID uint64
	Responses    chan<- *socket.Response
}

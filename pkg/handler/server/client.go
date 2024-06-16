package server

import "github.com/mymmrac/hide-and-seek/pkg/api"

type Client struct {
	Responses chan<- *api.Msg
}

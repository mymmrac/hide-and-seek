package chttp

import (
	"context"

	"github.com/mymmrac/hide-and-seek/pkg/module/ws"
)

type Client interface {
	Do(ctx context.Context, request *Request) (*Response, error)
	WS(ctx context.Context, url string) (ws.Connector, error)
}

type Request struct {
	URL    string
	Method string
	Body   []byte
}

type Response struct {
	Status int
	Body   []byte
}

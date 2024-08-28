//go:build wasm

package chttp

import (
	"context"
	"errors"

	"github.com/mymmrac/hide-and-seek/pkg/module/ws"
)

var DefaultClient = nil

type jsClient struct{}

func (c *jsClient) Do(ctx context.Context, request *Request) (*Response, error) {
	return nil, errors.ErrUnsupported
}

func (c *jsClient) WS(ctx context.Context, url string) (ws.Connector, error) {
	return nil, errors.ErrUnsupported
}

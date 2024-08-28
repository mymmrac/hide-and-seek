//go:build !wasm

package chttp

import (
	"context"
	"errors"
	"fmt"

	"github.com/fasthttp/websocket"
	"github.com/gofiber/fiber/v2"

	"github.com/mymmrac/hide-and-seek/pkg/module/ws"
)

var DefaultClient = NewFiberClient()

type fiberClient struct {
	httpClient *fiber.Client
	wsDialer   *websocket.Dialer
}

func NewFiberClient() Client {
	return &fiberClient{
		httpClient: &fiber.Client{},
		wsDialer:   websocket.DefaultDialer,
	}
}

func (c *fiberClient) Do(ctx context.Context, request *Request) (*Response, error) {
	switch request.Method {
	case fiber.MethodPost:
		status, data, errs := c.httpClient.Post(request.URL).Body(request.Body).Bytes()
		if len(errs) != 0 {
			return nil, errors.Join(errs...)
		}
		return &Response{
			Status: status,
			Body:   data,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported method: %s", request.Method)
	}
}

func (c *fiberClient) WS(ctx context.Context, url string) (ws.Connector, error) {
	conn, _, err := c.wsDialer.DialContext(ctx, url, nil)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

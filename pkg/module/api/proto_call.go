package api

import (
	"context"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/mymmrac/hide-and-seek/pkg/module/chttp"
)

func ProtoCall[TResponse, TRequest any, TProtoRequest TProto[TRequest], TProtoResponse TProtoErr[TResponse]](
	ctx context.Context, client chttp.Client, url string, request TProtoRequest,
) (TProtoResponse, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	resp, err := client.Do(ctx, &chttp.Request{
		URL:    url,
		Method: "POST",
		Body:   data,
	})
	if err != nil {
		return nil, fmt.Errorf("do request: %w", err)
	}

	if resp.Status != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.Status)
	}

	var response TProtoResponse = new(TResponse)
	if err = proto.Unmarshal(resp.Body, response); err != nil {
		return nil, fmt.Errorf("unmarshal response: %w", err)
	}

	if err = response.ValidateAll(); err != nil {
		return nil, fmt.Errorf("validate response: %w", err)
	}

	respErr := response.GetError()
	if respErr != nil {
		return nil, fmt.Errorf("api error: %w", &ErrorMessage{err: respErr})
	}

	return response, nil
}

package api

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/proto"

	"github.com/mymmrac/hide-and-seek/pkg/api/communication"
)

func ProtoHandler[TRequest, TResponse any, TProtoRequest TProto[TRequest], TProtoResponse TProto[TResponse]](
	handler func(fCtx *fiber.Ctx, request TProtoRequest) (TProtoResponse, error),
) func(fCtx *fiber.Ctx) error {
	return func(fCtx *fiber.Ctx) error {
		var request TProtoRequest = new(TRequest)
		if err := proto.Unmarshal(fCtx.Body(), request); err != nil {
			return fmt.Errorf("unmarshal request: %w", err)
		}

		if err := request.ValidateAll(); err != nil {
			return fmt.Errorf("validate request: %w", err)
		}

		response, err := handler(fCtx, request)
		if err != nil {
			return fmt.Errorf("handle request: %w", err)
		}

		if err = response.ValidateAll(); err != nil {
			return fmt.Errorf("validate response: %w", err)
		}

		data, err := proto.Marshal(response)
		if err != nil {
			return fmt.Errorf("marshal response: %w", err)
		}

		if _, err = fCtx.Write(data); err != nil {
			return fmt.Errorf("write response: %w", err)
		}

		return fCtx.SendStatus(fiber.StatusOK)
	}
}

func ProtoCall[TResponse, TRequest any, TProtoRequest TProto[TRequest], TProtoResponse TProtoErr[TResponse]](
	client *fiber.Client, url string, request TProtoRequest,
) (TProtoResponse, error) {
	data, err := proto.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	status, responseData, errs := client.Post(url).Body(data).Bytes()
	if len(errs) > 0 {
		return nil, fmt.Errorf("do request: %w", errors.Join(errs...))
	}

	if status != fiber.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", status)
	}

	var response TProtoResponse = new(TResponse)
	if err = proto.Unmarshal(responseData, response); err != nil {
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

type TProto[T any] interface {
	*T
	proto.Message
	ValidateAll() error
}

type TProtoErr[T any] interface {
	TProto[T]
	GetError() *communication.Error
}

type ErrorMessage struct {
	err *communication.Error
}

func (e *ErrorMessage) Error() string {
	return e.err.Message
}

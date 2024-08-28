//go:build !wasm

package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/protobuf/proto"
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

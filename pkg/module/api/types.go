package api

import (
	"google.golang.org/protobuf/proto"

	"github.com/mymmrac/hide-and-seek/pkg/api/communication"
)

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

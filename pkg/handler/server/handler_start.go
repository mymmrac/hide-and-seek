package server

import (
	"math/rand/v2"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/mymmrac/hide-and-seek/pkg/api/communication"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

func (s *Server) handlerStart(
	fCtx *fiber.Ctx, request *communication.Start_Request,
) (*communication.Start_Response, error) {
	ctx := fCtx.UserContext()
	logger.FromContext(ctx).Debugf("Start request: %+v", request)

	token := strconv.FormatUint(rand.Uint64(), 10)

	s.playerLock.Lock()
	s.players[token] = &Client{
		PlayerID:     rand.Uint64(),
		PlayerName:   request.Username,
		ConnectionID: 0,
		Responses:    nil,
	}
	s.playerLock.Unlock()

	return &communication.Start_Response{
		Type: &communication.Start_Response_Result_{
			Result: &communication.Start_Response_Result{
				Token: token,
			},
		},
	}, nil
}

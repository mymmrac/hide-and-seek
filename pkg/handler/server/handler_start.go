package server

import (
	"math/rand/v2"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/mymmrac/hide-and-seek/pkg/api/communication"
	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

func (s *Server) handlerStart(
	fCtx *fiber.Ctx, request *communication.Start_Request,
) (*communication.Start_Response, error) {
	ctx := fCtx.UserContext()
	logger.FromContext(ctx).Debugf("Start request: %+v", request)

	s.clients.Lock()
	defer s.clients.Unlock()
	players := s.clients.Raw()

	usernameUnique := true
	for _, player := range players {
		if player.PlayerName == request.Username {
			usernameUnique = false
			break
		}
	}
	if !usernameUnique {
		return &communication.Start_Response{
			Type: &communication.Start_Response_Error{
				Error: &communication.Error{
					Message: "Username is already taken",
				},
			},
		}, nil
	}

	token := strconv.FormatUint(rand.Uint64(), 10)
	players[token] = &Client{
		PlayerID:   rand.Uint64(),
		PlayerName: request.Username,
		StateLock:  sync.RWMutex{},
		State:      nil,
	}

	return &communication.Start_Response{
		Type: &communication.Start_Response_Result_{
			Result: &communication.Start_Response_Result{
				Token: token,
			},
		},
	}, nil
}

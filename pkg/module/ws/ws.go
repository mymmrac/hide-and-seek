package ws

import (
	"errors"
	"time"

	"github.com/fasthttp/websocket"

	"github.com/mymmrac/hide-and-seek/pkg/module/logger"
)

type Connector interface {
	ReadMessage() (messageType int, p []byte, err error)
	WriteMessage(messageType int, data []byte) error
	WriteControl(messageType int, data []byte, deadline time.Time) error
}

func ReadMessage(log *logger.Logger, conn Connector) ([]byte, bool) {
	if conn == nil {
		log.Errorf("Nil connection")
		return nil, false
	}

	_, data, err := conn.ReadMessage()
	if err != nil {
		if isExpectedCloseError(err) || errors.Is(err, websocket.ErrNilConn) {
			return nil, false
		}

		log.Warnf("Read message: %s", err)
		return nil, false
	}

	return data, true
}

func WriteMessage(log *logger.Logger, conn Connector, data []byte) bool {
	if conn == nil {
		log.Errorf("Nil connection")
		return false
	}

	if err := conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
		if !errors.Is(err, websocket.ErrCloseSent) && !errors.Is(err, websocket.ErrNilConn) {
			log.Warnf("Write message [%x]: %s", data, err)
		}
		return false
	}

	return true
}

func WriteCloseMessage(conn Connector) {
	_ = conn.WriteControl(websocket.CloseMessage, nil, time.Now().Add(time.Second))
}

var expectedCodes = []int{
	websocket.CloseNoStatusReceived,
	websocket.CloseGoingAway,
	websocket.CloseNormalClosure,
	websocket.CloseAbnormalClosure,
}

func isExpectedCloseError(err error) bool {
	var closeError *websocket.CloseError
	if errors.As(err, &closeError) {
		for _, code := range expectedCodes {
			if closeError.Code == code {
				return true
			}
		}
	}
	return false
}

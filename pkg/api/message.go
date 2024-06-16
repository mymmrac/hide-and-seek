package api

import (
	"bytes"
	"encoding/gob"

	"github.com/mymmrac/hide-and-seek/pkg/module/space"
)

type Msg struct {
	From uint64
	Pos  space.Vec2F
}

func (m *Msg) Marshal() ([]byte, error) {
	buf := bytes.NewBuffer(nil)
	if err := gob.NewEncoder(buf).Encode(m); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (m *Msg) Unmarshal(data []byte) error {
	if err := gob.NewDecoder(bytes.NewReader(data)).Decode(m); err != nil {
		return err
	}
	return nil
}

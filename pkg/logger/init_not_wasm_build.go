//go:build !wasm

package logger

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func init() {
	defaults := log.DefaultStyles()
	log.SetStyles(&log.Styles{
		Timestamp: defaults.Timestamp,
		Caller:    defaults.Caller,
		Prefix:    defaults.Prefix,
		Message:   defaults.Message,
		Key:       defaults.Key,
		Value:     defaults.Value,
		Separator: defaults.Separator,
		Levels: map[log.Level]lipgloss.Style{
			log.DebugLevel: defaults.Levels[log.DebugLevel].UnsetMaxWidth(),
			log.InfoLevel:  defaults.Levels[log.InfoLevel].UnsetMaxWidth(),
			log.WarnLevel:  defaults.Levels[log.WarnLevel].UnsetMaxWidth(),
			log.ErrorLevel: defaults.Levels[log.ErrorLevel].UnsetMaxWidth(),
			log.FatalLevel: defaults.Levels[log.FatalLevel].UnsetMaxWidth(),
		},
		Keys:   defaults.Keys,
		Values: defaults.Values,
	})
}

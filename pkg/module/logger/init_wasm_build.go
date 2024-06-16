//go:build wasm

package logger

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

func init() {
	log.SetStyles(&log.Styles{
		Timestamp: lipgloss.NewStyle(),
		Caller:    lipgloss.NewStyle(),
		Prefix:    lipgloss.NewStyle(),
		Message:   lipgloss.NewStyle(),
		Key:       lipgloss.NewStyle(),
		Value:     lipgloss.NewStyle(),
		Separator: lipgloss.NewStyle(),
		Levels: map[log.Level]lipgloss.Style{
			log.DebugLevel: lipgloss.NewStyle().SetString(strings.ToUpper(log.DebugLevel.String())),
			log.InfoLevel:  lipgloss.NewStyle().SetString(strings.ToUpper(log.InfoLevel.String())),
			log.WarnLevel:  lipgloss.NewStyle().SetString(strings.ToUpper(log.WarnLevel.String())),
			log.ErrorLevel: lipgloss.NewStyle().SetString(strings.ToUpper(log.ErrorLevel.String())),
			log.FatalLevel: lipgloss.NewStyle().SetString(strings.ToUpper(log.FatalLevel.String())),
		},
		Keys:   nil,
		Values: nil,
	})
}

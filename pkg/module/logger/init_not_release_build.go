//go:build !release

package logger

import "github.com/charmbracelet/log"

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
}

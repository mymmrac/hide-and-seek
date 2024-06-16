//go:build release

package logger

import "github.com/charmbracelet/log"

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(false)
}

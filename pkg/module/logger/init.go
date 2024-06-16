package logger

import "github.com/charmbracelet/log"

func init() {
	log.SetReportTimestamp(true)
	log.SetTimeFormat("2006.01.02 15:04:05")
}

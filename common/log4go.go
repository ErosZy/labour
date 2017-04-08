package common

import (
	"fmt"
)

const (
	LOG_DEBUG = iota
	LOG_WARNING
	LOG_FATAL
	LOG_INFO
)

func Logger(logType int, msg string) {
	switch logType {
	case LOG_DEBUG:
		fmt.Println("[DEBUG] " + msg)
	case LOG_WARNING:
		fmt.Println("[WANRING] " + msg)
	case LOG_FATAL:
		fmt.Println("[FATAL] " + msg)
	case LOG_INFO:
		fmt.Println("[INFO] " + msg)
	}
}

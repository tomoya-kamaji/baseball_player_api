package config

import "github.com/tomoya_kamaji/go-pkg/src/util/logging"

var l logging.Logger

// InitLogger is initialize logging
func InitLogger() {
	if l != nil {
		return
	}

	l = logging.NewStackDriverLogger("BaseballApi", logging.GetUserAgent)
}

// GetLogger is get logger
func GetLogger() logging.Logger {
	return l
}

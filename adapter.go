// Package golog provides a logger that writes to a github.com/kataras/golog log.
package golog

import (
	"context"

	"github.com/jackc/pgx/v5/tracelog"
	"github.com/kataras/golog"
)

// Logger supports golog integration with the new pgx's tracelogs.
type Logger struct {
	l *golog.Logger
}

// NewLogger returns a new logger which implements the pgx's tracelog's logger interface.
func NewLogger(l *golog.Logger) *Logger {
	return &Logger{l: l}
}

// Log implements the tracelog logger interface.
func (l *Logger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	logger := l.l

	switch level {
	case tracelog.LogLevelTrace:
		if len(data) > 0 {
			data["PGX_LOG_LEVEL"] = level
		} else {
			data = map[string]interface{}{"PGX_LOG_LEVEL": level}
		}

		logger.Debug(msg, golog.Fields(data))
	case tracelog.LogLevelDebug:
		logger.Debug(msg)
	case tracelog.LogLevelInfo:
		logger.Info(msg)
	case tracelog.LogLevelWarn:
		logger.Warn(msg)
	case tracelog.LogLevelError:
		logger.Error(msg)
	default:
		if len(data) > 0 {
			data["INVALID_PGX_LOG_LEVEL"] = level
		} else {
			data = map[string]interface{}{"INVALID_PGX_LOG_LEVEL": level}
		}
		logger.Error(msg, golog.Fields(data))
	}
}

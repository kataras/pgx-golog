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
//
// Example Code:
//
//	import (
//		"context"
//
//		"github.com/jackc/pgx/v5/pgxpool"
//		"github.com/jackc/pgx/v5/tracelog"
//		"github.com/kataras/golog"
//		pgxgolog "github.com/kataras/pgx-golog"
//	)
//
//	func main() {
//		golog.SetLevel("debug")
//
//		logger := pgxgolog.NewLogger(golog.Default)
//		tracer := &tracelog.TraceLog{
//			Logger:   logger,
//			LogLevel: tracelog.LogLevelTrace,
//		}
//
//		connString := "postgres://postgres:admin!123@localhost:5432/test_db?sslmode=disable&search_path=public"
//		connConfig, err := pgxpool.ParseConfig(connString)
//		if err != nil {
//			panic(err)
//		}
//		connConfig.ConnConfig.Tracer = tracer
//
//		pool, err := pgxpool.NewWithConfig(context.Background(), connConfig)
//		if err != nil {
//			panic(err)
//		}
//
//		// [...]
//	}
func NewLogger(l *golog.Logger) *Logger {
	return &Logger{l: l}
}

// Log implements the tracelog logger interface.
func (l *Logger) Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{}) {
	logger := l.l

	switch level {
	case tracelog.LogLevelTrace:
		if data == nil {
			data = make(map[string]interface{})
		}

		data["PGX_LOG_LEVEL"] = level

		logger.Debug(msg, golog.Fields(data))
	case tracelog.LogLevelDebug:
		logger.Debug(msg, golog.Fields(data))
	case tracelog.LogLevelInfo:
		logger.Info(msg, golog.Fields(data))
	case tracelog.LogLevelWarn:
		logger.Warn(msg, golog.Fields(data))
	case tracelog.LogLevelError:
		logger.Error(msg, golog.Fields(data))
	default:
		if data == nil {
			data = make(map[string]interface{})
		}

		data["INVALID_PGX_LOG_LEVEL"] = level

		logger.Error(msg, golog.Fields(data))
	}
}

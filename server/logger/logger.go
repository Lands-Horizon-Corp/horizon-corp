package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	return &Logger{logger}, nil
}

func (l *Logger) Close() {
	l.Sync()
}

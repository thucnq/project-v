package trace

import (
	"project-v/pkg/l"
)

type Log struct {
	l l.Logger
}

func NewTraceLog(l l.Logger) *Log {
	return &Log{l}
}
func (receiver Log) Debug(msg string) {
	receiver.l.Info(msg)
}

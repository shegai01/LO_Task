package logger

import (
	"log/slog"
)

type LogMessage struct {
	Message string
}
type ChanLogs struct {
	log chan LogMessage
}

func NewAsyncLog(buffer int) *ChanLogs {
	l := &ChanLogs{
		log: make(chan LogMessage, buffer),
	}
	go func() {
		for entry := range l.log {
			slog.Info(entry.Message)
		}
	}()

	return l
}
func (l *ChanLogs) Info(msg string) {
	select {
	case l.log <- LogMessage{Message: msg}:
	default:
	}
}
func (l *ChanLogs) Close() { close(l.log) }

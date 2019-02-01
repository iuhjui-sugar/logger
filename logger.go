package logger

import (
	"sync"

	"github.com/keystone-coin/logger/define"
	"github.com/keystone-coin/logger/format"
	"github.com/keystone-coin/logger/handler"
)

type Logger struct {
	handler *handler.SwapHandler
	mu      sync.Mutex
}

func NewLogger() *Logger {
	l := &Logger{}
	l.makeHandler()
	return l
}

func (l *Logger) makeHandler() {
	f := format.NewTerminalFormat()
	std := handler.NewStdoutHandler(f)
	swap := handler.NewSwapHandler(std)
	l.handler = swap
	return
}

func (l *Logger) write(ie define.IEntity) {
	l.handler.Log(ie)
}

func (l *Logger) SetHandler(other define.Handler) {
	l.handler.Swap(other)
}

func (l *Logger) Infof(template string, data map[string]string) {
	entity := define.NewEntity(define.INFO)
	entity.SetTemplate(template)
	entity.SetData(data)
	l.write(entity)
	return
}

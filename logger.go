package logger;

import (
    "sync"

    "github.com/keystone-coin/logger/define"
    "github.com/keystone-coin/logger/handler"
    "github.com/keystone-coin/logger/format"
)


type Logger struct {
    handler *handler.SwapHandler
    entitys map[string]define.IEntity
    mu    sync.Mutex
}

func NewLogger() *Logger {
    l := &Logger{}
    l.makeHandler();
    l.makeEntitys();
    return l;
}

func (l *Logger) makeHandler(){
    f    := format.NewTerminalFormat();
    std  := handler.NewStdoutHandler(f);
    swap := handler.NewSwapHandler(std);
    l.handler = swap;
    return;
}

func (l *Logger) makeEntitys() {
    l.entitys = map[string]define.IEntity{};
    return;
}

func (l *Logger) write(ie define.IEntity) {
    l.handler.Log(ie);
}

func (l *Logger) SetHandler(other define.Handler) {
	l.handler.Swap(other)
}

func (l *Logger) Bind(name string,ie define.IEntity){
    l.mu.Lock()
    defer l.mu.Unlock();
    l.entitys[name] = ie;
}

func (l *Logger) Emit(name string,data map[string]string){
    ie,exists := l.entitys[name];
    if exists == false {
        return
    }
    l.write(ie.Derive(data));
    return;
}

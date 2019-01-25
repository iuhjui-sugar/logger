package handler;

import (
    "os"

    "github.com/keystone-coin/logger/define"
)

type StdoutHandler struct{
    inner define.Handler
}

func NewStdoutHandler(f define.Format) *StdoutHandler{
    h:= &StdoutHandler{}
    h.inner = NewStreamHandler(os.Stdout,f);
    return h;
}

func (h *StdoutHandler) Log(ie define.IEntity) error {
	return h.inner.Log(ie);
}

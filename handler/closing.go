package handler;

import (
    "io"

    "github.com/keystone-coin/logger/define"
)

type ClosingHandler struct {
	w     io.WriteCloser
    inner define.Handler
}

func NewClosingHandler(w io.WriteCloser,inner define.Handler) *ClosingHandler  {
    h := &ClosingHandler{};
    h.w = w;
    h.inner = inner;
    return h;
}

func (h *ClosingHandler) Log(ie define.IEntity) error {
	return h.inner.Log(ie);
}

func (h *ClosingHandler) Close() error {
	return h.w.Close()
}

package handler;

import (
    "sync"

    "github.com/keystone-coin/logger/define"
)

type SafeHandler struct {
	inner define.Handler
	mu    sync.Mutex
}

func NewSafeHandler(inner define.Handler) *SafeHandler {
    h := &SafeHandler{};
    h.inner = inner;
    return h;
}

func (h *SafeHandler) Log(ie define.IEntity) error {
	h.mu.Lock()
	err := h.inner.Log(ie)
	h.mu.Unlock()
    if err != nil {
        return err
    }
    return nil;
}

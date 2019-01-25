package handler;

import (
    "sync"

    "github.com/keystone-coin/logger/define"
)

type SwapHandler struct {
    mu    sync.RWMutex
	inner define.Handler
}

func NewSwapHandler(inner define.Handler) *SwapHandler{
    h := &SwapHandler{}
    h.inner = inner;
    return h;
}

func (h *SwapHandler) Log(ie define.IEntity) error {
	h.mu.RLock()
	err := h.inner.Log(ie)
	h.mu.RUnlock()
    if err != nil {
        return err
    }
    return nil;
}

func (h *SwapHandler) Swap(other define.Handler) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.inner = other;
    return;
}

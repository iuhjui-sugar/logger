package handler;

import (
    "github.com/keystone-coin/logger/define"
)

type MultiHandler struct {
    handlers []define.Handler
}

func NewMultiHandler(handlers ...define.Handler) *MultiHandler  {
    h := &MultiHandler{};
    h.handlers = handlers;
    return h;
}

func (h *MultiHandler) Log(ie define.IEntity) error {
    for _,handler := range h.handlers {
        handler.Log(ie)
    }
    return nil
}

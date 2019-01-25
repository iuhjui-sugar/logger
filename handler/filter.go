package handler;

import (
    "github.com/keystone-coin/logger/define"
)

type FilterFunc func(ie define.IEntity) (skip bool)

type FilterHandler struct{
    inner define.Handler
    fn    FilterFunc
}

func NewFilterHandler(fn FilterFunc,inner define.Handler) *FilterHandler {
    h := &FilterHandler{};
    h.fn    = fn;
    h.inner = inner;
    return h;
}

func (h *FilterHandler) Log(ie define.IEntity) error {
	if h.fn(ie) == false {
		return h.inner.Log(ie);
	}
	return nil
}

package handler;

import (
    "github.com/keystone-coin/logger/define"
)

type BufferHandler struct {
    inner define.Handler
	recs  chan define.IEntity
}

func NewBufferHandler(size int,inner define.Handler) *BufferHandler{
    h := &BufferHandler{};
    h.inner = inner;
    h.recs  = make(chan define.IEntity,size)
    go h.run()
    return h;
}

func (h *BufferHandler) Log(ie define.IEntity) error {
	h.recs <- ie
	return nil
}

func (h *BufferHandler) run() {
	for ie := range h.recs {
		h.inner.Log(ie)
	}
}

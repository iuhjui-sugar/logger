package handler;

import (
    "github.com/keystone-coin/logger/define"
)

type LevelHandler struct {
    inner define.Handler
}

func NewLevelHandler(max define.Level,inner define.Handler) *LevelHandler {
    fn := func(ie define.IEntity) (skip bool){
        return ie.GetLevel() > max;
    }
    fh := NewFilterHandler(fn,inner);
    lh := &LevelHandler{};
    lh.inner = fh;
    return lh;
}

func (h *LevelHandler) Log(ie define.IEntity) error {
	return h.inner.Log(ie);
}

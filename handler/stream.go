package handler;

import (
	"io"

	"github.com/keystone-coin/logger/define"
)

type StreamHandler struct {
	w  io.Writer
	f  define.Format
}

func NewStreamHandler(w io.Writer, f define.Format) *StreamHandler {
	h := &StreamHandler{}
	h.w = w;
	h.f = f;
	return h;
}

func (h *StreamHandler) Log(ie define.IEntity) error {
	bytes := h.f.Format(ie);
	_, err := h.w.Write(bytes);
	if err != nil {
		return err
	}
	return nil;
}

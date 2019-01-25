package handler;

import (
    "os"

    "github.com/keystone-coin/logger/define"
)

type FileHandler struct{
    file  *os.File
    inner define.Handler
}

func NewFileHandler(path string, f define.Format) (*FileHandler, error) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
    sh := NewStreamHandler(file,f);
    ch := NewClosingHandler(file,sh);
    fh := &FileHandler{}
    fh.file  = file;
    fh.inner = ch;
	return fh, nil
}

func (h *FileHandler) Log(ie define.IEntity) error {
	return h.inner.Log(ie);
}

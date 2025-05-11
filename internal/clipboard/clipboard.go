package clipboard

import (
	"io"

	"golang.design/x/clipboard"
)

type C struct{}

var _ io.WriteCloser = (*C)(nil)

func New() *C {
	return &C{}
}

func Init() error {
	return clipboard.Init()
}

func (c *C) Write(p []byte) (n int, err error) {
	clipboard.Write(clipboard.FmtImage, p)
	return len(p), nil
}

func (c *C) Close() error {
	return nil
}

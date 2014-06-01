package writer

import (
	"bytes"
	"github.com/gopherjs/gopherjs/js"
	"io"
)

type Writer struct {
	buf *bytes.Buffer
	out js.Object
}

func New(out js.Object) *Writer {
	ret := new(Writer)
	ret.out = out
	ret.buf = new(bytes.Buffer)
	return ret
}

func (w *Writer) flush() {
	w.out.Call("println", w.buf.String())
	w.buf.Reset()
}

func (w *Writer) Write(buf []byte) (n int, e error) {
	for _, b := range buf {
		if b == byte('\n') {
			w.flush()
		} else if b >= 32 && b <= 126 {
			w.buf.WriteByte(b)
		} else {
			w.buf.WriteByte(byte('?'))
		}
	}

	return len(buf), nil
}

func (w *Writer) Close() error {
	if w.buf.Len() > 0 {
		w.flush()
	}
	return nil
}

var _ io.WriteCloser = new(Writer)

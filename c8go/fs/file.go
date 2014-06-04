package fs

import (
	"bytes"
	"io"
)

type File struct {
	perm  uint32
	bytes []byte
}

func NewFile(perm uint32) *File {
	ret := new(File)
	ret.perm = perm
	ret.bytes = make([]byte, 0)
	return ret
}

func NewStringFile(perm uint32, s string) *File {
	ret := NewFile(perm)
	ret.Set([]byte(s))
	return ret
}

func (f *File) Perm() uint32 { return f.perm }

func (f *File) Clone() *File {
	ret := new(File)
	ret.perm = f.perm

	writer := new(bytes.Buffer)
	io.Copy(writer, bytes.NewBuffer(f.bytes))
	ret.bytes = writer.Bytes()

	return ret
}

func (f *File) Reader() io.Reader {
	return bytes.NewBuffer(f.bytes)
}

func (f *File) Set(bytes []byte) {
	f.bytes = make([]byte, len(bytes))
	copy(f.bytes, bytes)
}

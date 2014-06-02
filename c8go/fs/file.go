package fs

import (
	"bytes"
)

type File struct {
	perm    uint32
	content *byte.Buffer
}

func NewFile(perm uint32) *File {
	ret := new(File)
	ret.perm = perm
	ret.content = new(byte.Buffer)
	return ret
}

func (f *File) Perm() uint32 { return f.perm }

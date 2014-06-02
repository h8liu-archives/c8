package fs

import (
	"bytes"
)

type File struct {
	perm uint32
	*bytes.Buffer
}

func NewFile(perm uint32) *File {
	ret := new(File)
	ret.perm = perm
	ret.Buffer = new(bytes.Buffer)
	return ret
}

func (f *File) Perm() uint32 { return f.perm }

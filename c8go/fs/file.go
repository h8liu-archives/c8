package fs

type File struct {
	perm    uint32
	content []byte
}

func (f *File) Perm() uint32 { return f.perm }

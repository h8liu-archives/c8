package fs

type Dir struct {
	perm uint32
	subs map[string]Node
}

func (d *Dir) Perm() uint32 { return d.perm }

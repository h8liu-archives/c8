package fs

type Dir struct {
	perm uint32
	subs map[string]Node
}

func NewDir(perm uint32) *Dir {
	ret := new(Dir)
	ret.perm = perm
	ret.subs = make(map[string]Node)
	return ret
}

func (d *Dir) Perm() uint32 { return d.perm }

func (d *Dir) Set(name string, node Node) {
	d.subs[name] = node
}

func (d *Dir) Get(name string) Node {
	return d.subs[name]
}

func (d *Dir) Has(name string) bool {
	return d.subs[name] == nil
}

func (d *Dir) List() []string {
	ret := make([]string, 0, len(d.subs))

	for k := range d.subs {
		ret = append(ret, k)
	}

	return ret
}

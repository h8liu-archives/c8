package fs

import (
	"path/filepath"
	"strings"
)

type FileSys struct {
	root *Dir
}

func NewFileSys() *FileSys {
	ret := new(FileSys)
	ret.root = NewDir(0)

	ret.buildSample()

	return ret
}

func (fs *FileSys) buildSample() {
	root := fs.root

	d := NewDir(0)
	d.Set("ls", NewFile(0))
	d.Set("mkdir", NewFile(0))
	d.Set("rm", NewFile(0))
	root.Set("bin", d)

	d = NewDir(0)
	f := NewFile(0)
	f.Write([]byte("# Readme file\n\n"))
	f.Write([]byte("Just a test file\n"))
	d.Set("readme", f)
	// d.Set("readme", NewStringFile(0, "# Readme file\n\nJust a test file\n"))

	home := NewDir(0)
	home.Set("h8liu", d)

	root.Set("home", home)
}

func (fs *FileSys) GetOrCreateDir(path string) *Dir {
	return fs.GetOrCreate(path, false).(*Dir)
}

func (fs *FileSys) GetOrCreateFile(path string) *File {
	return fs.GetOrCreate(path, true).(*File)
}

func (fs *FileSys) GetOrCreate(path string, file bool) Node {
	part, pnode := fs.GetLast(path)
	n := fs.GetPart(pnode, part)
	if n != nil {
		return n
	}
	if pnode != nil {
		if d, isDir := pnode.(*Dir); isDir {
			var node Node
			if file {
				node = NewFile(0)
			} else {
				node = NewDir(0)
			}
			d.Set(part, node)
			return d.Get(part)
		}
	}
	panic("no such file or directory") // can't resolve till last part
}

func (fs *FileSys) GetLast(path string) (string, Node) {
	if !filepath.IsAbs(path) {
		panic("has to be absolute path")
	}

	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return "", fs.root
	}

	if parts[0] != "" {
		panic("bug")
	}
	parts = parts[1:]

	var node Node = fs.root
	for _, p := range parts[:len(parts)-1] { // stop at the last part
		if p == "" {
			continue
		}
		node = fs.GetPart(node, p)
		if node == nil {
			return "", nil
		}
	}
	return parts[len(parts)-1], node
}

func (fs *FileSys) GetPart(node Node, part string) Node {
	if node == nil {
		return nil
	}
	d, isDir := node.(*Dir)
	if !isDir {
		return nil
	}
	return d.Get(part)
}

func (fs *FileSys) Get(path string) Node {
	// first resolve to last part, then get the last one
	// dividing into two steps are for sharing code
	// with GetOrCreate
	part, pnode := fs.GetLast(path)
	if part == "" { // has no parent!, return root
		return fs.root
	}
	return fs.GetPart(pnode, part)
}

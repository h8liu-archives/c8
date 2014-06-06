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

func (fs *FileSys) Get(path string) Node {
	if !filepath.IsAbs(path) {
		panic("has to be absolute path")
	}

	parts := strings.Split(path, "/")
	if parts[0] != "" {
		panic("bug")
	}
	parts = parts[1:]
	if parts[len(parts)-1] == "" {
		parts = parts[:len(parts)-1]
	}

	var node Node = fs.root
	for _, p := range parts {
		if p == "" {
			continue
		}

		d, isDir := node.(*Dir)
		if !isDir {
			return nil
		}

		node = d.Get(p)
		if node == nil {
			return nil // not found
		}
	}

	return node
}

package shell

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/h8liu/c8/c8go/fs"
)

func rm(args []string, out io.Writer) int {
	if len(args) != 2 {
		fmt.Fprintf(out, "rm needs an arg\n")
		return -1
	}

	rel := args[1]
	path := filepath.Join(Pwd, rel)
	
	if strings.HasPrefix(Pwd, path) {
		fmt.Fprintf(out, "cannot remove %q under %q\n", path, Pwd)
		return -1
	}

	dir, name := filepath.Split(path)
	node := fileSys.Get(dir)
	
	d, okay := node.(*fs.Dir)
	if !okay {
		fmt.Fprintf(out, "error: directory not exists\n")
		return -1
	}

	target := d.Get(name)
	if target == nil {
		fmt.Fprintf(out, "error: target not exists\n")
		return -1
	}

	tdir, isDir := target.(*fs.Dir)
	if isDir && !tdir.IsEmpty() {
		fmt.Fprintf(out, "error: directory not empty\n")
		return -1
	}

	d.Set(name, nil)
	return 0
}

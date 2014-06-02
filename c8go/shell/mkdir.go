package shell

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/h8liu/c8/c8go/fs"
)

func mkdir(args []string, out io.Writer) int {
	if len(args) != 2 {
		fmt.Fprintf(out, "mkdir needs an arg\n")
		return -1
	}

	rel := args[1]
	path := filepath.Join(Pwd, rel)
	dir, name := filepath.Split(path)

	node := fileSys.Get(dir)
	d, okay := node.(*fs.Dir)
	if !okay {
		fmt.Fprintf(out, "error: target directory does not exist\n")
		return -1
	}

	if !fs.IsValid(name) {
		fmt.Fprintf(out, "error: invalid directory name\n", name)
		return -1
	}

	if d.Get(name) != nil {
		fmt.Fprintf(out, "error: %q already exists\n", name)
		return -1
	}

	newDir := fs.NewDir(0)
	d.Set(name, newDir)

	return 0
}

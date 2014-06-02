package shell

import (
	"fmt"
	"io"

	"github.com/h8liu/c8/c8go/fs"
)

func mkdir(args []string, out io.Writer) int {
	if len(args) != 2 {
		fmt.Fprintf(out, "mkdir needs an arg\n")
		return -1
	}

	node := fileSys.Get(Pwd)
	dir, okay := node.(*fs.Dir)
	if !okay {
		fmt.Fprintf(out, "error: current working directory does not exist\n")
		return -1
	}

	name := args[1]

	if dir.Get(args[1]) != nil {
		fmt.Fprintf(out, "%q already exists\n", name)
		return -1
	}

	newDir := fs.NewDir(0)
	dir.Set(name, newDir)

	return 0
}

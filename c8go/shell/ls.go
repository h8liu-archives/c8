package shell

import (
	"fmt"
	"io"

	"github.com/h8liu/c8/c8go/fs"
)

func ls(args []string, out io.Writer) int {
	if len(args) != 1 {
		fmt.Fprintf(out, "ls with args not implemented yet\n")
		return -1
	}

	node := fileSys.Get(Pwd)
	dir, okay := node.(*fs.Dir)
	if !okay {
		fmt.Fprintf(out, "error: current working directory does not exist\n")
		return -1
	}

	lst := dir.List()
	for i, p := range lst {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, p)
	}
	fmt.Fprintln(out)

	return 0
}

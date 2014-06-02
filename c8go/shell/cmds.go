package shell

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/h8liu/c8/c8go/fs"
)

var fileSys = fs.NewFileSys()

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

func pwd(args []string, out io.Writer) int {
	fmt.Fprintln(out, Pwd)
	return 0
}

func cd(args []string, out io.Writer) int {
	if len(args) >= 3 {
		fmt.Fprintln(out, "cd taks at most one arg")
		return -1
	}

	if len(args) <= 1 {
		Pwd = "/"
		return 0
	}

	rel := args[1]

	pwd := filepath.Join(Pwd, rel)
	node := fileSys.Get(pwd)
	if node == nil {
		fmt.Fprintln(out, "directory not found")
		return -1
	}

	_, isDir := node.(*fs.Dir)
	if !isDir {
		fmt.Fprintln(out, "target is not a directory")
		return -1
	}

	Pwd = pwd
	return 0
}

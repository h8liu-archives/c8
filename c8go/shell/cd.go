package shell

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/h8liu/c8/c8go/fs"
)

func cd(args []string, out io.Writer) int {
	if len(args) >= 3 {
		fmt.Fprintln(out, "error: cd takes at most one arg")
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
		fmt.Fprintln(out, "error: directory not found or name invalid")
		return -1
	}

	_, isDir := node.(*fs.Dir)
	if !isDir {
		fmt.Fprintln(out, "error: target is not a directory")
		return -1
	}

	Pwd = pwd
	return 0
}

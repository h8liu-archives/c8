package shell

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/h8liu/c8/c8go/fs"
)

func cat(args []string, out io.Writer) int {
	if len(args) != 2 {
		fmt.Fprintf(out, "cat takes exactly 1 arg\n")
		return -1
	}

	path := filepath.Join(Pwd, args[1])
	node := fileSys.Get(path)
	file, isFile := node.(*fs.File)
	if !isFile {
		fmt.Fprintf(out, "error: file does not exist\n")
		return -1
	}

	reader := file.Reader()
	io.Copy(out, reader)
	return 0
}

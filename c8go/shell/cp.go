package shell

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/h8liu/c8/c8go/fs"
)

func cp(args []string, out io.Writer) int {
	if len(args) != 3 {
		fmt.Fprintf(out, "cp needs 2 args\n")
		return -1
	}

	from := filepath.Join(Pwd, args[1])
	to := filepath.Join(Pwd, args[2])

	fromDir, name := filepath.Split(from)
	if name == "" {
		fmt.Fprintf(out, "cannot move root\n")
		return -1
	}

	toNode := fileSys.Get(to)
	_, isDir := toNode.(*fs.Dir)
	if isDir {
		to = filepath.Join(to, name)
	}

	if strings.HasPrefix(Pwd, to) {
		fmt.Fprintf(out, "cannot move to %q under %q\n", to, Pwd)
		return -1
	}

	if from == to {
		fmt.Fprintf(out, "this copy is a noop\n")
		return -1
	}

	node := fileSys.Get(fromDir)
	d, okay := node.(*fs.Dir)
	if !okay {
		fmt.Fprintf(out, "error: directory not exists\n")
		return -1
	}

	target := d.Get(name)
	if target == nil {
		fmt.Fprintf(out, "error: copy target not exists\n")
		return -1
	}

	targetFile, isFile := target.(*fs.File)
	if !isFile {
		fmt.Fprintf(out, "error: we can only copy simple files now\n")
		return -1
	}

	check := fileSys.Get(to)
	if check != nil {
		_, okay = check.(*fs.Dir)
		if okay {
			fmt.Fprintf(out, "error: cannot overwrite a directory\n")
			return -1
		}
	}

	dir, rename := filepath.Split(to)
	dest := fileSys.Get(dir)
	destDir, okay := dest.(*fs.Dir)
	if !okay {
		fmt.Fprintf(out, "error: destination not exists\n")
		return -1
	}

	other := targetFile.Clone()
	destDir.Set(rename, other)

	return 0
}

package shell

import (
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/h8liu/c8/c8go/fs"
)

func mv(args []string, out io.Writer) int {
	if len(args) != 3 {
		fmt.Fprintf(out, "mv needs 2 args\n")
		return -1
	}

	from := filepath.Join(Pwd, args[1])
	to := filepath.Join(Pwd, args[2])

	if strings.HasPrefix(Pwd, from) {
		fmt.Fprintf(out, "cannot move %q under %q\n", from, Pwd)
		return -1
	}

	if strings.HasPrefix(Pwd, to) {
		fmt.Fprintf(out, "cannot move to %q under %q\n", from, Pwd)
		return -1
	}

	if from == to {
		fmt.Fprintf(out, "this move is a noop\n")
		return -1
	}

	dir, name := filepath.Split(from)
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

	dest := fileSys.Get(to)
	destDir, okay := dest.(*fs.Dir)
	if okay {
		d.Set(name, nil)
		destDir.Set(name, target)
		return 0
	}

	dir, rename := filepath.Split(to)
	dest = fileSys.Get(dir)
	destDir, okay = dest.(*fs.Dir)
	if !okay {
		fmt.Fprintf(out, "error: destination not exists\n")
		return -1
	}

	d.Set(name, nil)
	destDir.Set(rename, target)

	return 0
}

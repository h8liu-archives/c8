package shell

import (
	"fmt"
	"io"
	"path/filepath"

	"github.com/h8liu/c8/c8go/fs"
)

var Pwd string = "/"
var fileSys = fs.NewFileSys()

const (
	OutRedirect = ">"
	InRedirect = "<"
)

func System(args []string, out io.Writer) int {
	if len(args) == 0 {
		return 0
	}

	cmd := args[0]

	entry := builtin[cmd]
	if entry == nil {
		fmt.Fprintf(out, "command %q not found\n", cmd)
		return -1
	}

	/*
	  node := fileSys.GetOrCreate("/home/h8liu/test", true)
	  if node == nil {
			fmt.Fprintf(out, "cannot create")
			return -1
	  }
	*/

	var outfile *fs.File = nil
	var i int
	var arg string
	for i, arg = range args {
		if arg == OutRedirect {
			if i == len(args)-1 {
				fmt.Fprintf(out, "missing output redirection file\n")
				return -1
			}
			path := filepath.Join(Pwd, args[i+1])
			outfile = fileSys.GetOrCreateFile(path)
			outfile.Clear()
			break
		}
	}
	if outfile != nil {
		return entry(args[:i], outfile)
	} else {
		return entry(args, out)
	}
}

type EntryFunc func(args []string, out io.Writer) int

var builtin = map[string]EntryFunc{
	"ls":    ls,
	"pwd":   pwd,
	"mkdir": mkdir,
	"rm":    rm,
	"cp":    cp,
	"cd":    cd,
	"mv":    mv,
	"cat":   cat,
	"echo":  echo,
	"help":  help,
}

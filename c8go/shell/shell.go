package shell

import (
	"fmt"
	"io"
)

var Pwd string = "/"

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

	return entry(args, out)
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

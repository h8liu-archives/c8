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

	for i, a := range args {
		fmt.Fprintf(out, "%d: %s\n", i, a)
	}

	return 0
}

type EntryFunc func(args []string, out io.Writer)

var builtin = map[string]EntryFunc{
	"ls":    ls,
	"mkdir": mkdir,
	"rm":    rm,
	"cp":    cp,
	"cd":    cd,
	"mv":    mv,
	"cat":   cat,
	"echo":  echo,
}

package shell

import (
	"fmt"
	"io"
)

func echo(args []string, out io.Writer) int {
	args = args[1:]
	for i, a := range args {
		if i > 0 {
			fmt.Fprintf(out, " ")
		}
		fmt.Fprint(out, a)
	}
	return 0
}

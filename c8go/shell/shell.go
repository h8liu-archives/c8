package shell

import (
	"fmt"
	"io"
)

func System(args []string, out io.Writer) int {
	if len(args) == 0 {
		return 0
	}

	for i, a := range args {
		fmt.Fprintf(out, "%d: %s\n", i, a)
	}

	return 0
}

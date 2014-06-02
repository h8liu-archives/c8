package shell

import (
	"fmt"
	"io"
)

func pwd(args []string, out io.Writer) int {
	fmt.Fprintln(out, Pwd)
	return 0
}

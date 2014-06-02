package shell

import (
	"fmt"
	"io"
)

var (
	mkdir = notImpl
	rm    = notImpl
	cp    = notImpl
	mv    = notImpl
	cat   = notImpl
	echo  = notImpl
	help  = notImpl
)

func notImpl(args []string, out io.Writer) int {
	cmd := args[0]
	fmt.Fprintf(out, "command %q not implemented yet\n", cmd)
	return -1
}

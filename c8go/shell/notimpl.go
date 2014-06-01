package shell

import (
	"fmt"
	"io"
)

var (
	ls    = notImpl
	mkdir = notImpl
	rm    = notImpl
	cp    = notImpl
	cd    = notImpl
	mv    = notImpl
	cat   = notImpl
	echo  = notImpl
)

func notImpl(args []string, out io.Writer) {
	cmd := args[0]
	fmt.Fprintf(out, "command %q not implemented yet", cmd)
}

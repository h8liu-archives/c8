package shell

import (
	"fmt"
	"io"
	"strings"
)

var helpStr = strings.TrimSpace(`
supported built-in commands:
	ls     list the content of the current working directory
	pwd    print the current working directory
	mkdir  make a new directory
	rm     remove a file or an empty directory
	cp     copy a file
	mv     move a file
	cat    print the content of a file
	echo   print the args, separated by spaces
	help   print this help message
`)

func help(args []string, out io.Writer) int {
	fmt.Fprintf(out, helpStr)
	return 0
}

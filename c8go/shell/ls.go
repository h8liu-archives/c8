package shell

import (
	"fmt"
	"io"

	"github.com/h8liu/c8/c8go/fs"
)

func ls(args []string, out io.Writer) int {
  var items []string = []string{}
  var long bool
  for _, arg := range(args[1:]) {
    if arg == "-l" {
      long = true
    } else {
      items = append(items, arg)
    }
  }
  if len(items) == 0 {
    items = []string{Pwd}
  }
  for it, item:= range(items) {
    if len(items) > 1 {
		  fmt.Fprintln(out, item + ":")
    }
    if item[0] != '/' {
      item = Pwd + item
    }
    node := fileSys.Get(item)
    dir, okay := node.(*fs.Dir)
    if !okay {
      fmt.Fprintf(out, "error: current working directory does not exist\n")
      return -1
    }
    lst := dir.List()
    for i, p := range lst {
      if !long && i > 0 {
        fmt.Fprint(out, " ")
      }
      fmt.Fprint(out, p)
      if long {
        fmt.Fprintln(out)
      }
    }
    if len(lst) > 0 {
      if !long {
        fmt.Fprintln(out)
      }
      if len(items) > 1  && it != len(items) -1 {
        fmt.Fprintln(out)
      }
    }
  }

	return 0
}

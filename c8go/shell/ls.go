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
    full_item := item
    if item[0] != '/' {
      if Pwd[len(Pwd)-1] == '/' {
        full_item = Pwd + item
      } else {
        full_item = Pwd + "/" + item
      }
    }
    node := fileSys.Get(full_item)
    if node == nil {
      fmt.Fprintf(out, item + ": no such file or directory\n")
    } else {
      dir, okay := node.(*fs.Dir)
      if okay { // item is a directory, list its content
        lsdir(dir, long, out)
      } else { // item is a file, print the name
        fmt.Fprintln(out, item)
      }
    }
    if len(items) > 1  && it != len(items) -1 {
      fmt.Fprintln(out)
    }
  }

	return 0
}

func lsdir(dir *fs.Dir, long bool, out io.Writer) {
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
  if !long {
    fmt.Fprintln(out)
  }
}

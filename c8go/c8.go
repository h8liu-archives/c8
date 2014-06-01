package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
	"strings"

	"github.com/h8liu/c8/c8go/shell"
	"github.com/h8liu/c8/c8go/writer"
)

func main() {
	global := js.Global
	if global == nil {
		fmt.Println("not intended to run in a real OS")
		return
	}

	js.Global.Set("c8go", map[string]interface{}{
		"launch": Launch,
	})
}

var pout js.Object

func Println(s string) {
	pout.Call("println", s)
}

func Printf(f string, args ...interface{}) {
	s := fmt.Sprintf(f, args...)
	Println(s)
}

func SetOut(out js.Object) { pout = out }

func Launch(s string, out js.Object) {
	// SetOut(out)

	w := writer.New(out)
	shell.System(strings.Fields(s), w)
	w.Close()
}

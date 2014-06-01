package main

import (
	"github.com/gopherjs/gopherjs/js"
	"fmt"
	"strings"

	// "github.com/h8liu/c8/c8go/fs
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

func Println(out js.Object, s string) {
	out.Call("println", s)
}

func Printf(out js.Object, f string, args... interface{}) {
	s := fmt.Sprintf(f, args...)
	Println(out, s)
}

func Launch(s string, out js.Object) {
	fields := strings.Fields(s)
	for _, f := range fields {
		Println(out, f)
	}
}

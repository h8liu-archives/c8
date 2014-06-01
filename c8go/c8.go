package main

import (
	"github.com/gopherjs/gopherjs/js"
	"fmt"

	"github.com/h8liu/c8/fs"
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

func Launch(s string, out js.Object) {
	out.Call("println", "you typed: "+s)
	out.Call("println", fs.Hello())
}

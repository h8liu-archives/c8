package main

import (
	"fmt"
	"github.com/gopherjs/gopherjs/js"
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
	out.Call("println", "c8go says you typed: "+s)
}

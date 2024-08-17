// file libjs.go

// +build js

package SDK

import (
	"syscall/js"
)

var c chan string

func init() {
	c = make(chan string)
	js.Global().Set("read", js.FuncOf(reader))
}

func exit() {
	js.Global().Get("dostoy").Call("setPrompt", ">")
}

func println(text string) {
	js.Global().Get("dostoy").Call("println", text)
}

func read() string {
	c = make(chan string)
	return <-c
}

func reader(this js.Value, inputs []js.Value) interface{} {
	if len(inputs) > 0 {
		c <- inputs[0].String()

	} else {
		c <- ""
	}
	return nil
}
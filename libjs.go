// file libjs.go

// +build js

package SDK

import (
	"syscall/js"
	"os"
)

var c chan string

func init() {
	c = make(chan string)
	js.Global().Set("read", js.FuncOf(reader))
}

func Exit() {
	js.Global().Get("dostoy").Call("setPrompt", ">")
	os.Exit(3)
}

func Println(text string) {
	js.Global().Get("dostoy").Call("println", text)
}

func Print(text string) {
	js.Global().Get("dostoy").Call("print", text)
}

func Read() string {
	c = make(chan string)
	return <-c
}

func reader(this js.Value, inputs []js.Value) interface{} {
	if len(inputs) > 0 {
		c <- inputs[0].String()

	} else {
		c <- ""
	}
	return ""
}

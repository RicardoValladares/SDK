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

	/* INICIO AGREGADO */
	js.Global().Get("dostoy").Call("setPrompt", "")
	js.Global().Get("dostoy").Call("setShell", "false")
	js.Global().Get("dostoy").Call("println", "")
	/* FIN GREGADO */
}

func Exit() {
	/* INICIO AGREGADO */
	js.Global().Get("dostoy").Call("println", "")
	js.Global().Get("dostoy").Call("setPrompt", ">")
	js.Global().Get("dostoy").Call("setShell", "true")
	/* FIN GREGADO */
	os.Exit(3)
}

func Println(text string) {
	js.Global().Get("dostoy").Call("println", text)
}

func Print(text string) {
	js.Global().Get("dostoy").Call("print", text)
}

func Read() string {
	/* INICIO AGREGADO */
	js.Global().Get("dostoy").Call("setShell", "true")
	/* FIN GREGADO */
	
	c = make(chan string)
	
	/* INICIO AGREGADO */
	js.Global().Get("dostoy").Call("setShell", "false")
	/* FIN GREGADO */
	
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

// file libjs.go

// +build js

package SDK

import (
	"syscall/js"
	"os"
	"strings"
)

var c chan string
var r int

func init() {
	c = make(chan string)
	js.Global().Set("read", js.FuncOf(reader))

	/* INICIO AGREGADO */
	r = 0
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
	if r == 1 {
		js.Global().Get("dostoy").Call("print", ">")
	}
	/* FIN GREGADO */
	os.Exit(3)
}

func Println(text string) {

	/* INICIO AGREGADO */
	if strings.Contains(text,"\n"){
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			js.Global().Get("dostoy").Call("println", line)
		}
		return
	}
	/* FIN GREGADO */
	
	
	js.Global().Get("dostoy").Call("println", text)
}

func Print(text string) {
	
	/* INICIO AGREGADO */
	if strings.Contains(text,"\n"){
		lines := strings.Split(text, "\n")
		for index, line := range lines {
			if index == (len(lines)-1){
				js.Global().Get("dostoy").Call("print", line)
			} else {
				js.Global().Get("dostoy").Call("println", line)
			}
		}
		return
	}
	/* FIN GREGADO */
	
	js.Global().Get("dostoy").Call("print", text)
}

func Read() string {
	/* INICIO AGREGADO */
	r = 1
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

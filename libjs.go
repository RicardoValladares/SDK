// file libjs.go

// +build js

package SDK

import (
	"syscall/js"
	"os"
	"strings"
	qr "github.com/skip2/go-qrcode"
)

var c chan string
var r int

func init() {
	c = make(chan string)
	js.Global().Set("read", js.FuncOf(reader))
	r = 0
	js.Global().Get("dostoy").Call("setPrompt", "")
	js.Global().Get("dostoy").Call("setShell", "false")
	js.Global().Get("dostoy").Call("println", "")
}

func Exit() {
	js.Global().Get("dostoy").Call("println", "")
	js.Global().Get("dostoy").Call("setPrompt", ">")
	js.Global().Get("dostoy").Call("setShell", "true")
	if r == 0 {
		js.Global().Get("dostoy").Call("print", ">")
	}
	os.Exit(3)
}

func Println(text string) {
	if strings.Contains(text,"\n"){
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			js.Global().Get("dostoy").Call("println", line)
		}
		return
	}	
	js.Global().Get("dostoy").Call("println", text)
}

func Print(text string) {
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
	js.Global().Get("dostoy").Call("print", text)
}

func Read() string {
	r = 1
	js.Global().Get("dostoy").Call("setShell", "true")
	c = make(chan string)
	js.Global().Get("dostoy").Call("setShell", "false")
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

func QR(text string) {
	c, err := qr.New(text, qr.Highest)
	if err != nil {
		return
	}
	for ir, row := range c.Bitmap() {
		for ic, cell := range row {
			if ic!=0 && ir!=0 && ic!=1 && ir!=1 && ic!=2 && ir!=2 && ic!=34 && ir!=34 && ic!=35 && ir!=35 && ic!=36 && ir!=36 {
				if cell {
					js.Global().Get("dostoy").Call("color", "0", "0")
					js.Global().Get("dostoy").Call("print", " ")
				} else {
					js.Global().Get("dostoy").Call("color", "15", "15")
					js.Global().Get("dostoy").Call("print", " ")
				}	
			}
		}
		if ir!=34 && ir!=35 && ir!=36 {
			js.Global().Get("dostoy").Call("println", " ")
		}
	}
}

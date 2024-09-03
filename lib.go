// file lib.go

// +build !js

package SDK

import (
	"fmt"
	"os"
	qr "github.com/skip2/go-qrcode"
)

func Exit() {
	os.Exit(3)
}

func Println(text string) {
	fmt.Println(text)
}

func Print(text string) {
	fmt.Print(text)
}

func Read() string {
	var c string
	fmt.Scanf("%s", &c)
	return c
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
					fmt.Printf("\033[40m  \033[0m")
				} else {
					fmt.Printf("\033[47m  \033[0m")
				}	
			}
		}
		if ir!=34 && ir!=35 && ir!=36 {
			fmt.Println()
		}
	}
}

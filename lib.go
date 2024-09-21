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

func Shell() {
	os.Exit(3)
}

func Startx(display string) {
	fmt.Println("GUI: "+display)
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
	image := c.Bitmap()
	for ir, row := range image {
		for ic, cell := range row {
			if ic!=0 && ir!=0 && ic!=1 && ir!=1 && ic!=2 && ir!=2 && ic!=(len(image)-3) && ir!=(len(row)-3) && ic!=(len(image)-2) && ir!=(len(row)-2) && ic!=(len(image)-1) &&  ir!=(len(row)-1) {
				if cell {
					fmt.Printf("\033[40m  \033[0m")
				} else {
					fmt.Printf("\033[47m  \033[0m")
				}	
			}
		}
		if ir!=(len(row)-3) && ir!=(len(row)-2) && ir!=(len(row)-1) {
			fmt.Println()
		}
	}
}

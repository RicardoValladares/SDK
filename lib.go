// file lib.go

// +build !js

package SDK

import (
	"fmt"
	"os"
)

func init() { }

func Exit() {
	os.Exit(3)
}

func Println(text string) {
	fmt.Println(text)
}

func Read() string {
	var c string
	fmt.Scanf("%s", &c)
	return c
}
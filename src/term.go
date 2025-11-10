package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
)

func GetTermSize() (int, int) {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error could not get terminal size: ", err)
		os.Exit(1)
	}

	return width, height
}

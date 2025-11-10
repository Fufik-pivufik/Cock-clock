package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	fmt.Print("\033[?25l")
	inter := make(chan os.Signal, 1)
	signal.Notify(inter, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-inter

		fmt.Print("\033[H\033[2J")
		fmt.Print("\033[?25h")
		os.Exit(0)
	}()

	defer func() {
		fmt.Print("\033[H\033[2J")
		fmt.Print("\033[?25h")
	}()

	for {
		fmt.Print("\033[H\033[2J")
		times, date, loc := ParseDate()

		PrintD(times[:5], date, loc)
		time.Sleep(time.Second * 1)
	}
}

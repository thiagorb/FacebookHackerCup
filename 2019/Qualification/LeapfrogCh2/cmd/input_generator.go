package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func generator() chan string {
	var it func() chan string
	it = func() chan string {
		ch := make(chan string)
		go func() {
			ch <- ""

			g := it()

			for true {
				c := <-g
				ch <- c + "."
				ch <- c + "B"
			}
		}()

		return ch
	}

	return it()
}

func main() {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGPIPE)
		_ = <-c
		os.Exit(0)
	}()

	g := generator()
	for true {
		c := <-g
		fmt.Print("A")
		fmt.Println(c)
	}
}

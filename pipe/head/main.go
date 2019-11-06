package main

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signal_chan := make(chan os.Signal, 1)
	signal.Notify(signal_chan, syscall.SIGPIPE)

	go func() {
		select {
		case s := <-signal_chan:
			_, err := fmt.Println("catch signal")
			if err != nil {
				if errors.Is(err, syscall.EPIPE) {
					fmt.Fprintln(os.Stderr, "catch broken pipe error")
				} else {
					fmt.Fprintln(os.Stderr, err)
				}
			}
			fmt.Fprintln(os.Stderr, s)
		}
	}()

	run()
}

func run() {
	for {
		fmt.Println("Wow!")
	}
}
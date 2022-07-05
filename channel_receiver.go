package main

import (
	"fmt"
	"time"
)

func startSender(name string) <-chan string {
	c := make(chan string)
	go func() {
		c <- (name + " hello")
		time.Sleep(time.Second)
	}()
	return c
}
func main() {
	sender := startSender("Phu")

	fmt.Println(<-sender)
}

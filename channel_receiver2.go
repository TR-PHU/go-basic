package main

import (
	"fmt"
	"time"
)

func startSender(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 5; i++ {
			c <- (name + " hello")
			time.Sleep(time.Second)
		}
	}()
	return c
}
func main() {
	phu := startSender("Phu")
	duong := startSender("Duong")

	for {
		select {
		case msgPhu := <-phu:
			fmt.Println(msgPhu)
		case msgDuong := <-duong:
			fmt.Println(msgDuong)
		}
	}
}

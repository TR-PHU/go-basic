package main

import (
	"fmt"
	"time"
)

func main() {
	count := 0

	for i := 1; i <= 5; i++ {
		go func() {
			for i := 1; i < 10000; i++ {
				count++
			}
		}()
	}

	time.Sleep(time.Second * 10)
	fmt.Println("Count: ", count) // Count alway < 50000
}

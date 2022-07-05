package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int, 40)

	wg.Add(2)
	go func(ch <-chan int) { // receivier
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) { // sender
		ch <- 42
		ch <- 2
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	count := 0
	lock := new(sync.RWMutex) // Init mutex

	for i := 1; i <= 5; i++ {
		go func() {
			for i := 1; i <= 10000; i++ {
				lock.Lock() // allow 1 routine access race condition
				count++
				lock.Unlock() // announce this routine complete increase count + 1
			}
		}()
	}

	time.Sleep(time.Second * 10)
	fmt.Println("Count: ", count) // Count : 50000
}

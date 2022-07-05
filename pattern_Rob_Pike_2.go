package main

import (
	"fmt"
	"math/rand"
	"time"
)

func query(url string) string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
	return url
}

func queryFirst(servers ...string) <-chan string {
	c := make(chan string)

	for _, serv := range servers {
		go func(s string) {
			c <- query(s)
		}(serv)
	}
	return c
}
func main() {
	result := queryFirst("server 1", "server 2", "server 3")
	fmt.Println(<-result)
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func fetchAPI(model string) string {
	time.Sleep(time.Duration(rand.Intn(1e3)) * time.Microsecond)
	return model
}
func main() {
	responseChan := make(chan string)
	var results []string

	go func() { responseChan <- fetchAPI("users") }()
	go func() { responseChan <- fetchAPI("categories") }()
	go func() { responseChan <- fetchAPI("products") }()

	for i := 1; i <= 3; i++ {
		results = append(results, <-responseChan)
	}

	fmt.Println(results)
}

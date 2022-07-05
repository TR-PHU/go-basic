package main

import "fmt"

func main() {
	randonmNumbers := []int{}
	for i := 1; i <= 10; i++ {
		randonmNumbers = append(randonmNumbers, i)
	}
	sum := 0
	for i := 0; i < len(randonmNumbers); i++ {
		sum += randonmNumbers[i] * randonmNumbers[i]
	}
	fmt.Printf("Total sum of square: %d", sum)
}

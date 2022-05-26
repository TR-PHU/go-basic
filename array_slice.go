package main

import "fmt"

func main() {
	// slice
	a := []int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
	a = append(a, 5, 6)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	// Array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)
	fmt.Println(len(b))
	fmt.Println(cap(b))
}

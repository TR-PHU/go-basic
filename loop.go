package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
	array := []int{2, 3, 4, 5}
	for index, value := range array {
		fmt.Println(index, value)
	}
	m := map[string]int{
		"Zeus":  20,
		"Duong": 17,
		"Mie":   20,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
	s := "Hello world!"
	for i, v := range s {
		fmt.Println(i, string(v))
	}
}

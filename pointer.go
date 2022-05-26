package main

import "fmt"

func main() {
	// var a int = 12
	// var b *int = &a
	// fmt.Println(a, b)  // Output 12 address of a
	// fmt.Println(a, *b) // Output 12 12
	// a = 24
	// fmt.Println(a, *b) // Out put 24 24
	// *b = 32
	// fmt.Println(a, *b)

	// m := []int{1, 2, 3} // slice
	// n := m
	// fmt.Println(m, n)
	// m[1] = 10
	// fmt.Println(m, n)

	// var a *myStruct
	// fmt.Println(a) // nil == null
	// a = new(myStruct)
	// fmt.Println(a)           // &{0}
	// (*a).number = 20         // access value pointer obj a
	// fmt.Println((*a).number) // get value pointer obj a

	var myMap = map[string]int{
		"Zeus": 20,
		"Mie":  20,
	}
	mapCopy := myMap
	myMap["Zeus"] = 21
	fmt.Println(myMap, mapCopy)
}

type myStruct struct {
	number int
}

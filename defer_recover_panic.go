package main

import "fmt"

func main() {
	// defer fmt.Println("Line 1")
	// defer fmt.Println("Line 2")
	// defer fmt.Println("Line 3")
	// Push to stack
	// Output Line 3 Line 2 Line 1
	fmt.Println("Start")
	panicker()
	fmt.Println("End")

}
func panicker() {
	fmt.Println("Create panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error :", err)
			// panic(err)
		}
	}()
	panic("This is a panic")
}

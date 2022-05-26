package main

import "fmt"

func main() {
	m := map[string]int{
		"Zeus": 19,
		"Mie":  20,
	}
	if age, isExist := m["Zeus"]; isExist {
		fmt.Println("Age of Zeus", age)
	}
	// Swtich case
	var i interface{} = "Hello world!" // Interface
	switch i.(type) {
	case int:
		fmt.Println("Kieu int")
	case float64:
		fmt.Println("Kieu float64")
	default:
		fmt.Println("Others")
	}
}

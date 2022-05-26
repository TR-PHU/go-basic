package main

import "fmt"

func main() {
	// a := 10
	// b := 200
	// // c := []int{3, 12, 90, -1, -2}
	// max := findMax(a, b)
	// fmt.Println(*max)
	// // fmt.Println(findSum((c)))
	// fmt.Println(findSum(1, 2, 3, 4, 5))
	// res, err := calDidive(4, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(res)
	g := greeter{
		greeting: "Hello",
		name:     "GO",
	}
	g.greet()
	fmt.Println(g.name)
}

func findMax(a, b int) *int {
	a = 100
	if a > b {
		return &a
	}
	return &b
}
func findSum(a ...int) (sum int) {
	for _, v := range a {
		sum += v
	}
	return
}
func calDidive(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() { // Method of struct
	fmt.Println(g.greeting, g.name)
	g.name = "Zerus"
}

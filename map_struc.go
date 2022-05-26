package main

import "fmt"

type Date struct {
	day, month, year int
}
type Product struct {
	name    string
	price   int
	expires Date
}

func main() {
	// Map

	studentNameAgeMap := make(map[string]int)
	studentNameAgeMap = map[string]int{
		"Phu":   20,
		"Thuy":  18,
		"Phuoc": 16,
	}
	delete(studentNameAgeMap, "Phu")
	fmt.Println(studentNameAgeMap)
	fmt.Println(studentNameAgeMap["Phu"])
	student := Student{
		number: 1,
		name:   "Phu chan",
		isMale: true,
		subjects: []string{
			"Math",
			"English",
			"Computer",
		},
	}
	fmt.Println(student.number)

	// Struct

	product := Product{
		name:  "Mi Hao Hao",
		price: 3500,
		expires: Date{
			day:   1,
			month: 1,
			year:  2023,
		},
	}
	fmt.Println(product)
}

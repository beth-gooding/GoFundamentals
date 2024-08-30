package main

import "fmt"

type person struct {
	lName  string
	age    int
	salary float64
}

func main() {
	fName := "Joe"
	grades := []int{100, 68, 23}
	states := map[string]string{"KY": "Kentucky", "VA": "Virginia", "WV": "West Virginia"}
	p := person{lName: "Smith", age: 444, salary: 12345678.90}

	fmt.Printf("fName value: %#v\n", fName)
	fmt.Printf("grades value: %#v\n", grades)
	fmt.Printf("states value: %#v\n", states)
	fmt.Printf("person value: %#v\n", p)

}

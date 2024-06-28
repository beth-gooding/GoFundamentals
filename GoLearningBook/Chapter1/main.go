package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	printStars()
	doctorsOfficeForm()
}

func printStars() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)

	fmt.Println(stars)
}

func doctorsOfficeForm() {
	var firstName string = "Hazy"
	var familyName string = "AI"
	var age int = 28
	var peanutAllergy bool = false

	fmt.Println(firstName)
	fmt.Println(familyName)
	fmt.Println(age)
	fmt.Println(peanutAllergy)
}

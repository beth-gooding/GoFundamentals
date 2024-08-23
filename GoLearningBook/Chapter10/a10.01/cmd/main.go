package main

import (
	"fmt"
	"strings"

	"a10.01/pkg/payroll"
)

var reviewRatings = make(map[string]interface{})

func init() {
	fmt.Println("Welcome to the employee pay and performance review")
	fmt.Println(strings.Repeat("+", 50))
}

func init() {
	fmt.Println("Initializing variables")
	reviewRatings["Work Quality"] = "Fair"
	reviewRatings["Teamwork"] = 2
	reviewRatings["Skills"] = "Poor"
}

func main() {
	d := payroll.Developer{Individual: payroll.Employee{FirstName: "Eric", LastName: "Davis", Id: 1}, HoursWorkedInYear: 1000, HourlyRate: 84.00, Review: reviewRatings}
	m := payroll.Manager{Individual: payroll.Employee{FirstName: "The", LastName: "Boss", Id: 2}, Salary: 150000, CommissionRate: 0.1}

	fmt.Printf("%v %v got a review rating of %f\n", d.Individual.FirstName, d.Individual.LastName, d.ReviewRating())
	payroll.PayDetails(d)
	payroll.PayDetails(m)
}

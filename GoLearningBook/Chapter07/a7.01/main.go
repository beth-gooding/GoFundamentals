package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}

type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func (d Developer) Pay() (string, float64) {
	yearlyPay := d.HourlyRate * d.HoursWorkedInYear
	name := d.Individual.FirstName + " " + d.Individual.LastName
	return name, yearlyPay
}

func (m Manager) Pay() (string, float64) {
	yearlyPay := m.Salary * (1 + m.CommissionRate)
	name := m.Individual.FirstName + " " + m.Individual.LastName

	return name, yearlyPay
}

func payDetails(p Payer) {
	fullName, yearlyPay := p.Pay()
	fmt.Println(fullName, "got paid", yearlyPay, "for the year")
}

func convertRating(s string) float64 {
	switch rating := s; rating {
	case "Excellent":
		return 5.00
	case "Good":
		return 4.00

	case "Fair":
		return 3.00

	case "Poor":
		return 2.00

	case "Unsatisfactory":
		return 1.00

	default:
		return 0.00
	}

}

func (d Developer) reviewRating() float64 {
	totalScore := 0.00
	for _, i := range d.Review {
		switch v := i.(type) {
		case string:
			totalScore += convertRating(v)
		case int:
			totalScore += float64(v)
		default:
			totalScore += 0
		}
	}

	return totalScore / float64(len(d.Review))
}

func main() {
	reviewRatings := make(map[string]interface{})
	reviewRatings["Work Quality"] = "Fair"
	reviewRatings["Teamwork"] = 2
	reviewRatings["Skills"] = "Poor"
	d := Developer{Individual: Employee{FirstName: "Eric", LastName: "Davis", Id: 1}, HoursWorkedInYear: 1000, HourlyRate: 84.00, Review: reviewRatings}
	m := Manager{Individual: Employee{FirstName: "The", LastName: "Boss", Id: 2}, Salary: 150000, CommissionRate: 0.1}

	fmt.Printf("%v %v got a review rating of %f\n", d.Individual.FirstName, d.Individual.LastName, d.reviewRating())
	payDetails(d)
	payDetails(m)
}

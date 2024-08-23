package payroll

import "fmt"

type Payer interface {
	Pay() (string, float64)
}

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

func PayDetails(p Payer) {
	fullName, yearlyPay := p.Pay()
	fmt.Println(fullName, "got paid", yearlyPay, "for the year")
}

package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

type directDeposit struct {
	lastName      string
	firstName     string
	bankName      string
	routingNumber int
	accountNumber int
}

func (d *directDeposit) validateRoutingNumber() error {
	if d.routingNumber < 100 {
		panic(ErrInvalidRoutingNumber)
	}
	return nil
}

func (d *directDeposit) validateLastName() error {
	// Not specific enough, should think about whitespace strings too
	if d.lastName == "" {
		return ErrInvalidLastName
	}
	return nil
}

func (d *directDeposit) report() {
	// Can improve using strings.repeat
	fmt.Println("*********************************************************************************************")
	fmt.Println("Last Name:", d.lastName)
	fmt.Println("First Name:", d.firstName)
	fmt.Println("Bank Name:", d.bankName)
	fmt.Println("Routing Number:", d.routingNumber)
	fmt.Println("Account Number:", d.accountNumber)
}

func main() {
	myDirectDeposit := directDeposit{lastName: "", firstName: "Abe", bankName: "XYZ Inc", routingNumber: 17, accountNumber: 1809}

	err := myDirectDeposit.validateLastName()
	if err != nil {
		fmt.Println(err)
	}

	err = myDirectDeposit.validateRoutingNumber()
	if err != nil {
		fmt.Println(err)
	}

	myDirectDeposit.report()
}

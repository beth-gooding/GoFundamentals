package main

import (
	"errors"
	"fmt"
)

func isApproved(interestRate float64, monthlyPayment float64, monthlyIncome float64) bool {
	if monthlyPayment > (monthlyIncome * interestRate) {
		return false
	}

	return true
}

func loanCalculator(applicant string, creditScore float64, monthlyIncome float64, loanAmount float64, loanTerm int) error {
	var interestRate float64

	if creditScore < 0 || monthlyIncome < 0 || loanAmount < 0 || loanTerm < 0 {
		return errors.New("Inputs cannot be negative")
	}

	if loanTerm%12 != 0 {
		return errors.New("Loan term must be divisible by 12 months")
	}

	if creditScore >= 450 {
		interestRate = 0.15

	} else {
		interestRate = 0.2
	}

	monthlyPayment := (loanAmount * (1 + interestRate)) / float64(loanTerm)
	approved := isApproved(interestRate, monthlyPayment, monthlyIncome)

	totalCost := (monthlyPayment * float64(loanTerm)) - loanAmount

	fmt.Println("Applicant", applicant)
	fmt.Println("------------")
	fmt.Println("Credit Score :", creditScore)
	fmt.Println("Income :", monthlyIncome)
	fmt.Println("Loan Amount :", loanAmount)
	fmt.Println("Loan Term :", loanTerm)
	fmt.Println("Monthly Payment :", monthlyPayment)
	fmt.Println("Rate :", interestRate)
	fmt.Println("Total Cost : ", totalCost)
	fmt.Println("Approved :", approved)
	fmt.Println("")

	return nil
}

func main() {
	loanCalculator("1", 500, 1000, 1000, 24)
	loanCalculator("2", 350, 1000, 10000, 12)
}

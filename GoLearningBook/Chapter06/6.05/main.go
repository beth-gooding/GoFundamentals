package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func payDay(hoursWorked, hourlyRate int) int {
	defer func() {
		if r := recover(); r != nil {
			if r == ErrHourlyRate {
				fmt.Printf("hourly rate: %d\nerr: %v\n\n", hourlyRate, r)
			}
			if r == ErrHoursWorked {
				fmt.Printf("hours worked: %d\nerr: %v\n\n", hoursWorked, r)
			}

			fmt.Printf("Pay was calculated based on: \nhours worked: %d\nhourly rate: %d\n", hoursWorked, hourlyRate)
		}
	}()
	if hourlyRate < 10 || hourlyRate > 75 {
		panic(ErrHourlyRate)
	}
	if hoursWorked < 0 || hoursWorked > 80 {
		panic(ErrHoursWorked)
	}

	if hoursWorked > 40 {
		hoursOver := hoursWorked - 40
		overTime := hoursOver * 2
		regularPay := 40 * hourlyRate
		return regularPay + overTime
	}
	return hoursWorked * hourlyRate
}

func main() {
	pay := payDay(100, 25)
	fmt.Println("Pay 1:", pay)
	pay = payDay(100, 200)
	fmt.Println("Pay 2:", pay)
	pay = payDay(60, 25)
	fmt.Println("Pay 3:", pay)
}

package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(day Weekday, hours int) {
	d.WorkWeek[day] = hours
}

func (d *Developer) HoursWorked() int {
	cumulativeHoursWorked := 0

	for _, value := range d.WorkWeek {
		cumulativeHoursWorked += value
	}
	return cumulativeHoursWorked
}

func nonLoggedHours() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func (d *Developer) PayDay() (int, bool) {
	hoursWorked := d.HoursWorked()
	if hoursWorked > 40 {
		overtimeHours := hoursWorked - 40
		overtimePay := overtimeHours * d.HourlyRate * 2
		totalPay := (40 * d.HourlyRate) + overtimePay
		return totalPay, true
	}
	totalPay := hoursWorked * d.HourlyRate
	return totalPay, false
}

func (d *Developer) PayDetails() {

	for day, v := range d.WorkWeek {
		switch day {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Println("Monday hours: ", v)
		case 2:
			fmt.Println("Tuesday hours: ", v)
		case 3:
			fmt.Println("Wednesday hours: ", v)
		case 4:
			fmt.Println("Thursday hours: ", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}

	totalPay, isOvertime := d.PayDay()
	fmt.Println()
	fmt.Println("Hours worked this week:", d.HoursWorked())
	fmt.Println("Pay for the week:", totalPay)
	fmt.Println("Is this overtime pay:", isOvertime)
}

func main() {
	myDeveloper := Developer{Individual: Employee{1, "Beth", "G"}, HourlyRate: 100}

	x := nonLoggedHours()
	fmt.Println("Hours worked so far today:", x(2))
	fmt.Println("Hours worked so far today:", x(3))
	fmt.Println("Hours worked so far today:", x(5))
	fmt.Println()

	myDeveloper.LogHours(1, 8)
	myDeveloper.LogHours(2, 10)
	myDeveloper.LogHours(3, 10)
	myDeveloper.LogHours(4, 10)
	myDeveloper.LogHours(5, 6)
	myDeveloper.LogHours(6, 8)

	myDeveloper.PayDetails()

}

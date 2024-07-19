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

func main() {
	myDeveloper := Developer{Individual: Employee{1, "Beth", "G"}, HourlyRate: 100}

	mondayHours := 6
	tuesdayHours := 10

	myDeveloper.LogHours(1, mondayHours)
	fmt.Println("Hours worked on Monday:", mondayHours)

	myDeveloper.LogHours(2, tuesdayHours)
	fmt.Println("Hours worked on Tuesday:", tuesdayHours)

	fmt.Println("Hours worked this week:", myDeveloper.HoursWorked())
}

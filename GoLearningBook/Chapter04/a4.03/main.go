package main

import "fmt"

func main() {
	daysOfWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	daysOfWeek = append(daysOfWeek[6:], daysOfWeek[0:6]...)

	fmt.Println(daysOfWeek)
}

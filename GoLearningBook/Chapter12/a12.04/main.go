package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("The current time:", now.Format(time.ANSIC))

	duration := ((6 * 60 * 60) + (6 * 60) + 6) * time.Second
	futureTime := now.Add(duration)

	fmt.Println("The time 6 hours, 6 minutes, and 6 seconds from now is:", futureTime.Format(time.ANSIC))

}

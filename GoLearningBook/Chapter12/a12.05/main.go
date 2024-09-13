package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("The local current time is:", now.Format(time.ANSIC))

	nyTimeZone, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	nyTime := now.In(nyTimeZone)
	fmt.Println("The NY current time is:", nyTime.Format(time.ANSIC))

	laTimeZone, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println(err)
	}
	laTime := now.In(laTimeZone)
	fmt.Println("The LA current time is:", laTime.Format(time.ANSIC))
}

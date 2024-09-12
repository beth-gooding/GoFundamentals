package main

import (
	"fmt"
	"time"
)

func timeDiff(timeZone string) (string, string) {
	current := time.Now()
	remoteZone, err := time.LoadLocation(timeZone)
	if err != nil {
		fmt.Println(err)
	}
	remoteTime := current.In(remoteZone)
	fmt.Println("The current time is:", current.Format(time.ANSIC))
	fmt.Println("The timezone:", timeZone, "time is:", remoteTime)
	return current.Format(time.ANSIC), remoteTime.Format(time.ANSIC)
}

func main() {
	fmt.Println(timeDiff("America/Los_Angeles"))
}

package main

import (
	"fmt"
	"strconv"
	"time"
)

func formatTime() string {
	current := time.Now()
	day := current.Day()
	month := current.Month()
	year := current.Year()
	hour := current.Hour()
	minute := current.Minute()
	seconds := current.Second()

	return strconv.Itoa(hour) + ":" + strconv.Itoa(minute) + ":" + strconv.Itoa(seconds) + " " + strconv.Itoa(year) + "/" + strconv.Itoa(int(month)) + "/" + strconv.Itoa(day)
}

func main() {
	fmt.Println(formatTime())
}

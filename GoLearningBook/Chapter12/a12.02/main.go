package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	date := time.Date(2024, time.September, 13, 15, 15, 0, 0, time.Local)
	year := strconv.Itoa(date.Year())
	month := strconv.Itoa(int(date.Month()))
	day := strconv.Itoa(date.Day())
	hour := strconv.Itoa(date.Hour())
	minute := strconv.Itoa(date.Minute())
	second := strconv.Itoa(date.Second())

	fmt.Println(hour + ":" + minute + ":" + second + " " + year + "/" + month + "/" + day)
}

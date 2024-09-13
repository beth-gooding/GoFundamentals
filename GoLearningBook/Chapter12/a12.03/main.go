package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	time.Sleep(2 * time.Second)
	later := time.Now()

	duration := later.Sub(now)

	fmt.Println("The duration was:", duration)
}

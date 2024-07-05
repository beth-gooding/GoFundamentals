package main

import "fmt"

func main() {
	mySlice := []string{"Good", "Good", "Bad", "Good", "Good"}

	mySlice = append(mySlice[:2], mySlice[3:]...)

	fmt.Println(mySlice)
}

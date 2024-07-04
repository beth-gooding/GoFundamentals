package main

import "fmt"

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	return arr
}

func main() {
	var myArr [10]int
	myArr = fillArray(myArr)

	fmt.Println(myArr)
}

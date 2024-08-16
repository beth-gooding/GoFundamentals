package main

import "fmt"

func findMinGeneric[Num int | float64](nums []Num) Num {
	if len(nums) == 0 {
		return -1
	}

	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func main() {
	intSlice := []int{1, 32, 5, 8, 10, 11}
	floatSlice := []float64{1.1, 32.1, 5.1, 8.1, 10.1, 11.1}

	minGenericInt := findMinGeneric(intSlice)
	fmt.Printf("%v is the minimum integer value\n", minGenericInt)
	minGenericFloat := findMinGeneric(floatSlice)
	fmt.Printf("%v is the minimum float value\n", minGenericFloat)

}

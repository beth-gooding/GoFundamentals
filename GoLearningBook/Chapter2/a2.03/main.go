package main

import "fmt"

func main() {
	nums := []int{3, 2, 14, 15, 7, 12, 6, 1}
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	fmt.Println(nums)

}

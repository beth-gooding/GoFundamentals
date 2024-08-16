package main

import "fmt"

func FindLargestRanchStock[K comparable, V int | float64](m map[K]V) K {
	var stock V
	var name K

	for k, v := range m {
		if v > stock {
			stock = v
			name = k
		}
	}
	return name
}

func main() {
	animalStock := map[string]int{
		"Chicken": 5,
		"Cattle":  20,
		"Horses":  4,
	}

	miscStock := map[string]float64{
		"Hay":        5.5,
		"Feed":       1.2,
		"Fertilizer": 4.5,
	}

	largestStockOnRangeInt := FindLargestRanchStock(animalStock)
	fmt.Printf("The largest stocked item on the ranch is %s\n", largestStockOnRangeInt)
	largestStockOnRangeFloat := FindLargestRanchStock(miscStock)
	fmt.Printf("The largest stocked item on the ranch is %s\n", largestStockOnRangeFloat)
}

package main

import "exercise10.01/pkg/shape"

func main() {
	t := shape.Triangle{Height: 2, Base: 3}
	r := shape.Rectangle{Height: 7, Width: 4}
	s := shape.Square{Side: 5}

	shape.PrintShapeDetails(t, r, s)
}

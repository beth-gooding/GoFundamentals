package main

import "fmt"

type Shape interface {
	Area() float64
	Name() string
}

type triangle struct {
	base   float64
	height float64
}

type rectangle struct {
	height float64
	width  float64
}

type square struct {
	side float64
}

func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t triangle) Name() string {
	return "triangle"
}

func (r rectangle) Area() float64 {
	return r.height * r.width
}

func (r rectangle) Name() string {
	return "rectangle"
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Name() string {
	return "square"
}

func printShapeDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is: %.2f\n", item.Name(), item.Area())
	}
}

func main() {
	t := triangle{height: 2, base: 3}
	r := rectangle{height: 7, width: 4}
	s := square{side: 5}

	printShapeDetails(t, r, s)
}

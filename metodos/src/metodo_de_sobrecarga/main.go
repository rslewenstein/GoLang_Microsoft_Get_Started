package main

import "fmt"

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

type coloredTriangle struct {
	triangle
	color string
}

func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}

func main() {
	t := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", t.size)
	fmt.Println("Perimeter (colored)", t.perimeter())
	fmt.Println("Perimeter (normal)", t.triangle.perimeter())
}

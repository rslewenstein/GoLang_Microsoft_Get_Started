package main

import "fmt"

type triangle struct {
	size int
}

// type square struct {
// 	size int
// }

func (t triangle) perimeter() int {
	return t.size * 3
}

// func (s square) perimeter() int {
// 	return s.size * 4
// }

func (t *triangle) doubleSize() {
	t.size *= 2
}

func main() {
	t := triangle{3}
	t.doubleSize()
	fmt.Println("Size: ", t.size)
	fmt.Println("Perimeter (square): ", t.perimeter())
}

package main

import "fmt"

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())
}


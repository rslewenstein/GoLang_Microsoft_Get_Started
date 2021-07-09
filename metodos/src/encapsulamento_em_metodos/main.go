package main

import (
	"fmt"

	"github.com/rslewenstein/GoLang_Microsoft_Get_Started/blob/main/metodos/src/encapsulamento_em_metodos/geometry.go"
)

func main() {
	t := geometry.Triangle{}
	t.SetSize(3)
	fmt.Println("Perimeter", t.Perimeter())
}

package main

import (
	"fmt"
	"math"
)

var name string = "Marcin"

//isCool := true // Non-declaration definition outside function body

func main() {
	age := 24
	ballzSize := 4.51
	isCool := true

	fmt.Println("Jaja", name, "age", age, "with ballz of", ballzSize, "in diameter.")
	fmt.Println("I am", "cool?", isCool)
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
}

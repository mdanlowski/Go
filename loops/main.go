package main

import "fmt"

func main() {
	// lognish
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// shorthand
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}

	// FizzBuzz
	for i := 1; i <= 100; i++ {
		if i%15 == 0 {
			fmt.Println("fizzbuzz", "\t", i)
		} else if i%5 == 0 {
			fmt.Println("buzz", "\t", i)
		} else if i%3 == 0 {
			fmt.Println("fizz", "\t", i)
		}
	}
}

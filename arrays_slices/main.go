package main

import "fmt"

func slays() []string {
	return []string{"Test", "test", "tesT"}
}

func main() {
	// declare
	var testArr [2]string

	// populate
	testArr[0] = "I am"
	testArr[1] = "Ballzeboss"

	// test
	fmt.Println(testArr)

	// decl and assign
	testSlice := []string{"In place", "Assignment", "Ayy"}
	fmt.Printf("%T\n", testArr)
	fmt.Printf("%T\n", testSlice)
	fmt.Printf("%T\n", slays())

	fmt.Println("range: ", testArr[0:1])
	fmt.Println("length: ", len(testSlice))
}

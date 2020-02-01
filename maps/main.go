package main

import "fmt"

func main() {
	fmt.Println("Oh hi Mark")

	// define a map
	emails := make(map[string]string)
	// assign kvs
	emails["Andrzej"] = "test@test1.com"
	emails["Jessica"] = "tset@test2.com"

	fmt.Println("map size:", len(emails))
	fmt.Println("deleting")
	delete(emails, "Andrzej")
	fmt.Println("new map size:", len(emails))

	// declare+assign
	declaredEmails := map[string]string{"Janusz": "janusz@ballz.net", "Bomba": "dzida@ddd.dd"}

	// RANGE
	for key, val := range declaredEmails {
		fmt.Print(key, " ")
		fmt.Println(val) // ==> declaredEmails[key]
	}

	ids := []int{1, 25, 64, 724}
	// loop thru ids using range
	// index, element - cool!
	for i, id := range ids {
		fmt.Printf("%d\tID: %d\n", i, id)
	}
	// not using index
	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}
}

package main

import "fmt"

func main() {
	// START OMIT
	type A struct {
		Name string	`json:"name"`
	}
	type B struct {
		Name string	`json:"fullName"`
	}

	a := A{"Gopher"}
	b := B(a) // valid since go 1.8 // HL

	// this still does not work, converting must be explicit: // HL
	// b = a // HL
	// END OMIT

	fmt.Printf("b %v\n", b)
}

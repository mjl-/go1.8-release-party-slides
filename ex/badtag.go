package main

func main() {
	// START OMIT
	type Person struct {
		FirstName string `json:"name",xml:"name"` // bad: comma // HL
		LastName string `json:"name"` // bad: duplicate "name"  // HL
	}
	// END OMIT
}

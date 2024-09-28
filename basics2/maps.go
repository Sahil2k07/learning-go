package basics2

import "fmt"

func maps() {
	// Maps in Go
	ages := make(map[string]int)
	ages["John"] = 37
	ages["Mary"] = 35

	// Insert
	ages["Robin"] = 19

	fmt.Println(ages)

	// Get an element
	fmt.Println(ages["Robin"])

	// Delete the element
	delete(ages, "Robin")

	// Check the element if its present
	elem, ok := ages["Robin"]

	fmt.Println(elem, ok)

	map2 := map[string]int{
		"John": 37,
		"Mary": 35,
	}

	fmt.Println(map2)

	// Nested Maps
	nestedMap := map[string]map[string]int{
		"user1": {
			"id":   1,
			"auth": 66,
		},
		"user2": {
			"id":   1,
			"auth": 44,
		},
	}

	fmt.Println(nestedMap)
}

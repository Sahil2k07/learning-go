package basics1

import "fmt"

func conditionals() {
	age := 18

	if age > 18 {
		fmt.Println("You are an adult")
	} else if age > 16 && age <= 18 {
		fmt.Println("You are getting there")
	} else {
		fmt.Println("You are a kid")
	}

	// The conditionals operators are all the same
}

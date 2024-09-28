package basics1

import "fmt"

func loops() {
	// For loop in Go

	for i := 0; i < 6; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	// There are no while loops in Go

	val := 0

	for val < 5 {
		fmt.Print(val, " ")

		val++
	}

	fmt.Println()

	fruits := []string{"Apple", "Banana", "Melons"}

	// Kinda works like a for-each loop
	for i, fruit := range fruits {
		fmt.Println(i, fruit)
	}
}

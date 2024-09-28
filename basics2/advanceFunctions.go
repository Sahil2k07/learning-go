package basics2

import "fmt"

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// Higher Order function in Go
func calculator(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

func createMultiplier(x int) func(int) int {
	return func(y int) int {
		return x * y
	}
}

func advanceFunctions() {
	a := 2

	b := 3

	fmt.Println(calculator(a, b, add))
	fmt.Println(calculator(a, b, multiply))

	double := createMultiplier(2)

	fmt.Println(double(2))

	triple := createMultiplier(3)

	fmt.Println(triple(3))
}

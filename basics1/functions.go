package basics1

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func doubleReturns(s1, s2 string, n1 int) (string, int) {
	return s1 + s2, n1
}

func implicitReturns() (x, y int) {
	return
	// The default value initialised is 0 so it returns 0 and can also do this

	// return x, y
}

func functions() {
	addResults := add(2, 3)

	addResults2 := add2(3, 4)

	fmt.Println(addResults, addResults2)

	concatString, nums := doubleReturns("shahil ", "sharma", 0)

	fmt.Println(concatString, nums)

	val1, _ := implicitReturns()
	// Since we do not need the second value we use a _ instead of a variable name
	// The compiler as well discards the returned variable

	fmt.Println(val1)

	// Anonymous Functions
	func() {
		fmt.Println("Inside the Anonymous function")
	}()
}

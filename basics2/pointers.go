package basics2

import "fmt"

func double(x int) int {
	return x * 2
}

// Pointers in functions
func doublePointer(x *int) {
	*x *= 2
}

func pointers() {
	num := 2

	// ptr has the address of num stord in the ptr variable

	// First syntax
	// var ptr *int = &num

	// Second syntax
	ptr := &num

	fmt.Println("The Pointer value of ptr is", ptr)

	fmt.Println("The Actual value of ptr is", *ptr)

	// using de-reference operator we directly change the value of num here
	*ptr = 4

	fmt.Println("The updated value of num is", num)

	a := 4

	fmt.Println("Without Pointer, calling double function, the value returned from the function is", double(a))

	fmt.Println("But the original value is still", a)

	doublePointer(&a)

	fmt.Println("After calling the function with pointer the value of a is", a)

}

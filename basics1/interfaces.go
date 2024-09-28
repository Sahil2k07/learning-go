package basics1

import "fmt"

type Calculator interface {
	add1() int
	multiply() int
}

type Numbers struct {
	num1 int
	num2 int
}

func (n Numbers) add() string {
	return fmt.Sprintf("The addition result is %v", n.num1+n.num2)
}

func (n Numbers) add1() int {
	return n.num1 + n.num2
}

func (n Numbers) multiply() int {
	return n.num1 * n.num2
}

func interfaces() {
	nums := Numbers{
		num1: 2,
		num2: 3,
	}

	nums.add()

	// calc := nums
	var calc Calculator = nums

	fmt.Println("The calc is ", calc)

	fmt.Println(calc.add1())

	fmt.Println(calc.multiply())
}

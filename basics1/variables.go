package basics1

import "fmt"

func variables() {
	var nums int = 21

	opt := 2
	// Type infered

	var name string = "shahil"

	var price float32 = 30000.00
	// Chosing a float32 or 64 is compulsory

	var isSolution bool = true

	//Always a better choice
	name2 := "sharma"

	fmt.Println(nums, opt, name, price, isSolution, name2)

	// Prints the type of the variable instead of the value
	fmt.Printf("%T", price)

	fmt.Println()

	//Can declare variable in a single line
	place, value := "Delhi", 22

	fmt.Println(place, value)

	// Thsi is allowed for the variables
	nums = 22
	place = "Mumbai"
	value = 23

	fmt.Println(nums, place, value)

	//But Const works as teh same in javascript
	const permanentValue = 23

	fmt.Println(permanentValue)

	msg := fmt.Sprintf("The name is %v and the name2 is %v ", name, name2)

	fmt.Println(msg)
}

package basics1

import "fmt"

func slices() {
	// Arrays in Go
	arr := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(arr)

	fmt.Println(arr[2])

	// Slices
	slice := arr[1:3] // Index 1 to 2 (3 excluded)

	fmt.Println(slice)

	// Make method
	// This initialises a slice with initial length 5 and maximum capcacity 10
	// What this means is there are five values 0 in it and we can append more 5 values i.e. till the length is 10
	mySlice1 := make([]int, 5, 10)

	// If we dont sepecify the capacity it defaults to length
	mySlice2 := make([]int, 5)

	mySlice1 = append(mySlice1, 2)

	mySlice2 = append(mySlice2, 5)

	fmt.Println(mySlice1, mySlice2)

	fmt.Println(mySlice1[5], mySlice2[5])

	// Slice Literals
	sliceLiteral := []string{"Go", "is", "Awesome"}

	fmt.Println(sliceLiteral)

	// Capacity and Length functions for slices
	fmt.Println(len(mySlice1), cap(mySlice1))

}

package basics3

import "fmt"

// An interface for types that can be printed
type Printer interface {
	printVals() string
}

// A generic function that works with any type that implements Printer
func PrintStringers[T Printer](p []T) {
	for _, v := range p {
		fmt.Println(v.printVals())
	}
}

// A type that implements Stringer
type Person struct {
	Name string
}

func (p Person) printVals() string {
	return p.Name
}

func generics() {
	people := []Person{{Name: "Alice"}, {Name: "Bob"}}
	PrintStringers(people)
}

package basics2

import "fmt"

// Closures in Go
func concatter() func(string) string {
	doc := ""
	return func(word string) string {
		doc += word + " "
		return doc
	}
}

func addValue() func(int) int {
	val := 0
	return func(num int) int {
		val += num
		return val
	}
}

func closures() {
	closure1 := concatter()

	closure1("Let's")
	closure1("insert")
	closure1("s")
	closure1("string")

	fmt.Println(closure1(""))

	closure2 := addValue()

	closure2(2)

	fmt.Println(closure2(0))
}

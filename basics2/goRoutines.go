package basics2

import (
	"fmt"
	"time"
)

func printHello() {
	fmt.Println("Hello From the Go routines")
}

func goRoutines() {
	go printHello()

	// Without this line, we wont be able to see the output from printHello as the function is being executed in a different core
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from the main function")
}

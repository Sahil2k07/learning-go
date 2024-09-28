package basics2

import (
	"fmt"
	"time"
)

// Only a string value can be sent to a channell and is a good practice
func sendMessage(c chan<- string) {
	c <- "Hello from channel"
}

func worker(id int, ch chan<- string) {
	time.Sleep(time.Duration(id) * time.Second)
	ch <- fmt.Sprintf("Worker %v finished the work", id)
}

func channels() {
	// Channels are used to interact with goRoutines
	msgChannel := make(chan string)

	go sendMessage(msgChannel)

	// Receives the data from a channel
	msg := <-msgChannel

	// Close the channel, that means that no more values will be sent or received
	close(msgChannel)

	fmt.Println(msg)

	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()

	// The select statement allows a goroutine to wait on multiple communication operations
	select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout!")
	}

	workerChannel := make(chan string)

	for i := 1; i <= 3; i++ {
		go worker(i, workerChannel)
	}

	for i := 1; i <= 3; i++ {
		msg := <-workerChannel
		fmt.Println(msg)
	}
}

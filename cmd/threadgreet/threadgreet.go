package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChannel chan bool) {
	fmt.Println("Hello!", phrase)
	doneChannel <- true // send our bool to indicate that we're done
}

func slowGreet(phrase string, doneChannel chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello Slow!", phrase)
	doneChannel <- true // send our bool to indicate that we're done

	// since this is the slowest we can close the channel, which will in turn tell the 'for' loop
	// below that things are complete (but this only works in a case like this because we know the
	// slowest goroutine.  If we don't close the channel the 'for' loop will throw an error.
	close(doneChannel)
}

func main() {
	// create a slice to hold 4 separate channels
	//dones := make([]chan bool, 4)
	//
	//dones[0] = make(chan bool)
	//go greet("Nice to meet you!", dones[0])
	//dones[1] = make(chan bool)
	//go greet("How are you?", dones[1])
	//dones[2] = make(chan bool)
	//go slowGreet("How ... are ... you ...?", dones[2])
	//dones[3] = make(chan bool)
	//go greet("I hope you're liking the course!", dones[3])

	// wait for all done channels to complete
	//for _, done := range dones {
	//	<-done
	//}

	// Alternative 1...
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	// iterate using the channel until the channel is closed.  We could change the 'for' loop to
	// get the value (bool) returned by the channel, but in this case we don't need it.
	for range done {

	}

}

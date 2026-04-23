package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
	close(doneChan) // this close() will only work when we know which func will takes the longest
}

func main() {
	// // creating channel
	// done := make(chan bool)
	// // greet("Nice to meet you!")
	// // greet("How are you?")
	// // slowGreet("How ... are ... you ...?")
	// // greet("I hope you're liking the course!")
	// go greet("Nice to meet you!", done)
	// go greet("How are you?", done)
	// go slowGreet("How ... are ... you ...?", done)
	// go greet("I hope you're liking the course!", done)

	// // <-done
	// // <-done
	// // <-done
	// // <-done

	/////////////////////////
	// dones := make([]chan bool, 4)
	// dones[0] = make(chan bool)
	// go greet("Nice to meet you!", dones[0])
	// dones[1] = make(chan bool)
	// go greet("How are you?", dones[1])
	// dones[2] = make(chan bool)
	// go slowGreet("How ... are ... you ...?", dones[2])
	// dones[3] = make(chan bool)
	// go greet("I hope you're liking the course!", dones[3])
	// for _, done := range dones {
	// 	<-done
	// }
	//////////////////////////
	done := make(chan bool)
	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)
	for range done {

	}
}

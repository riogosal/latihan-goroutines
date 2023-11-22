package main

import (
	"fmt"
	"time"
)

func FakeRequest(sequence int, abortSignal chan bool, doneCh chan bool) {
	select {
	case <-time.After(100 * time.Second):
		fmt.Println(sequence, "I waited 100s")
		doneCh <- true // mengirim
	case <-abortSignal:
		fmt.Println(sequence, "I was aborted")
		doneCh <- false // mengirim
	}
	return
}

func main() {
	doneCh := make(chan bool)
	abortSignal := make(chan bool)

	// go func() {
	// 	ch1 := make(chan bool)
	// 	defer close(ch1)

	// 	go func(channel chan bool) {
	// 		time.Sleep(5 * time.Second)
	// 		channel <- true
	// 	}(ch1)
	// }()

	go FakeRequest(1, abortSignal, doneCh)

	time.Sleep(2 * time.Second)
	abortSignal <- true

	doneSuccess := <-doneCh // mendengar sama dengan sync.Wait()
	if doneSuccess {
		fmt.Println("I'm done")
	} else {
		fmt.Println("Timeout aborted")
	}

	fmt.Println("I'm done")
}

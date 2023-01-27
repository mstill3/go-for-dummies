package main

import (
	"fmt"
	"time"
)

func myChannels() {

	// By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value
	messages := make(chan string)

	// goroutine anonymous function
	go func() { 
		messages <- "ping"
	}()

	msg := <-messages
	fmt.Println(msg)


	// Buffering
	// Buffered channels accept a limited number of values without a corresponding receiver for those values
	texts := make(chan string, 2)
	texts <- "hows it going?"
	texts <- "not much, hbu"
	fmt.Println(<-texts)
	fmt.Println(<-texts)


	// Synchronization
	worker := func (done chan bool) {
		fmt.Print("working...")
		time.Sleep(time.Second)
		fmt.Println("done")	   
		// Send a value to notify that we’re done 
		done <- true
	}

	done := make(chan bool, 1)
    	go worker(done)
	// Block until we receive a notification from the worker on the channel
	// If you removed the <- done line from this program, the program would exit before the worker even started
    	<-done


	// Directions
	// only accepts a channel for sending values
	ping := func (pings chan<- string, msg string) {
		pings <- msg
	}
	
	// accepts one channel for receives & a second for sends
	pong := func (pings <-chan string, pongs chan<- string) {
		msg := <-pings
		pongs <- msg
	}

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

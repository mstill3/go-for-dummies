package main

import (
	"fmt"
	"time"
)

// define whats present in a error interface
type error interface {
	Error() string
}

type MyError struct {
	time time.Time
	message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v: %s", e.time, e.message)
}

func run() error {
	return &MyError{time.Now(), "issue not working"}
}

func myErrors() {
	// error is a very common interface
	// A nil error denotes success; a non-nil error denotes failure
	// As with fmt.Stringer, the fmt package looks for the error interface when printing values
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

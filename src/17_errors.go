package main

import (
	"fmt"
	"time"
	"math"
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



type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	} else {
		return 0.0, ErrNegativeSqrt(x)
	}
}

func myErrors() {
	// error is a very common interface
	// A nil error denotes success; a non-nil error denotes failure
	// As with fmt.Stringer, the fmt package looks for the error interface when printing values
	if err := run(); err != nil {
		fmt.Println(err)
	}

	if val, err := sqrt(2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
	if val, err := sqrt(-2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(val)
	}
}

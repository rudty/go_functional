package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeTick(t *testing.T) {
	for range time.Tick(2 * time.Second) {
		fmt.Println("a")
		// fmt.Println(a)
	}
}

type myError struct {
}

func (myError) Error() string {
	return ""
}

func zeroNil() (int, error) {
	var a *myError = nil
	return 0, a
}

func zeroNilWrapper() (int, error) {
	v, err := zeroNil()
	return v, err
}

func TestZeroNil(t *testing.T) {
	v, err := zeroNilWrapper()
	fmt.Println(v, err == nil) // false
}

func TestPanic(t *testing.T) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)
		}
	}()
	panic(nil)
}

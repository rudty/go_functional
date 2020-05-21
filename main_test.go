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

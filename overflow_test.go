package main

import (
	"fmt"
	"testing"
)

const uint32Max = ^uint32(0)
const int32Max = int32((^uint32(0)) >> 1)

func TestOverFlow(t *testing.T) {
	fmt.Println(uint32Max)
	fmt.Println(int32Max)
	var a int32 = 1
	fmt.Println(int32Max + a)
}

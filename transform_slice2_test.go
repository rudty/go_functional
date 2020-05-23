package main

import (
	"fmt"
	"testing"
)

type intMappingFunction func(int) int

func intMapInternal(a []int, fn intMappingFunction) []int {
	r := make([]int, len(a))
	for i, e := range a {
		r[i] = fn(e)
	}
	return r
}

func intMap(a []int, fn ...intMappingFunction) ([]int, func(intMappingFunction) []int) {
	if len(fn) == 0 {
		return nil, func(f intMappingFunction) []int {
			return intMapInternal(a, f)
		}
	}
	return intMapInternal(a, fn[0]), nil
}

func getIntArray() []int {
	return []int{1, 2, 3, 4, 5}
	// randArray := make([]int, 10)
	// for i := 0; i < len(randArray); i++ {
	// 	randArray[i] = int(rand.Uint32())
	// }
	// return randArray
}

func TestIntMap(t *testing.T) {
	a := getIntArray()
	r, _ := intMap(a, func(e int) int {
		return e * 2
	})
	fmt.Println(r)

	_, f := intMap(a)
	r2 := f(func(e int) int {
		return e + 3
	})
	fmt.Println(r2)
}

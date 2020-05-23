package main

import (
	"fmt"
	"testing"
)

type intFunctor struct {
	a []int
}

func (f intFunctor) Map(fn func(int) int) intFunctor {
	for i, e := range f.a {
		f.a[i] = fn(e)
	}
	return f
}

func (f intFunctor) String() string {
	return fmt.Sprint(f.a)
}
func magical(a []int) intFunctor {
	r := intFunctor{a}
	return r
}

func doSomething(a int) int {
	return a * 2
}

func getIntSlice() []int {
	return []int{1, 2, 3, 4, 5}
	// randArray := make([]int, 10)
	// for i := 0; i < len(randArray); i++ {
	// 	randArray[i] = int(rand.Uint32())
	// }
	// return randArray
}
func TestSlice1(t *testing.T) {
	ints := getIntSlice()
	results := make([]int, len(ints))
	for i, elt := range ints {
		results[i] = doSomething(elt)
	}
}

func TestSlice2(t *testing.T) {
	ints := getIntSlice()
	results := magical(ints).Map(doSomething)
	fmt.Printf("%s\n", results)
}

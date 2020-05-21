package main_test

import (
	"fmt"
	"strconv"
	"testing"
)

type strFunc func(string) string

func compose(f strFunc, g strFunc) strFunc {
	return func(s string) string {
		return g(f(s))
	}
}
func TestCollection(t *testing.T) {
	numbers := []interface{}{
		1,
		5,
		3,
		2,
	}
	fmt.Print(numbers)
	// a := 1
	// b := a + 1.8 error
	// fmt.Print(b)

	a := 1
	b := "a" + strconv.Itoa(a)
	// strconv.Atoi(b)
	fmt.Println(b)

}

func TestCompose(t *testing.T) {
	f := func(s string) string {
		return s + "a"
	}
	s := func(s string) string {
		return s + "b"
	}
	r := compose(f, s)("q")
	fmt.Println(r)

}

type intCurryingFunction func(a ...int) (int, interface{})

func myCurryAdd3(a ...int) (int, interface{}) {
	if len(a) == 3 {
		return a[0] + a[1] + a[2], nil
	}
	if len(a) == 1 {
		return 0, func(b ...int) (int, interface{}) {
			return myCurryAdd3(a[0], b[0])
		}
	}
	return 0, func(b ...int) (int, interface{}) {
		return myCurryAdd3(a[0], a[1], b[0])
	}
}

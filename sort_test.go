package main

import (
	"fmt"
	"math/rand"
	"testing"
)

func getRandArray() []int {
	randArray := make([]int, 10)
	for i := 0; i < len(randArray); i++ {
		randArray[i] = int(rand.Uint32())
	}
	return randArray
}

func mySort(arr []int, left, right int) {
	if left < right {
		c := arr[left]

		less := make([]int, 0)
		greater := make([]int, 0)
		same := []int{c}

		for i := left + 1; i <= right; i++ {
			if c == arr[i] {
				same = append(same, arr[i])
			} else if c > arr[i] {
				less = append(less, arr[i])
			} else {
				greater = append(greater, arr[i])
			}
		}

		mySort(less, 0, len(less)-1)
		mySort(greater, 0, len(greater)-1)

		for i := 0; i < len(less); i++ {
			arr[left] = less[i]
			left++
		}
		for i := 0; i < len(same); i++ {
			arr[left] = same[i]
			left++
		}
		for i := 0; i < len(greater); i++ {
			arr[left] = greater[i]
			left++
		}
	}
}

func TestSort(t *testing.T) {
	arr := getRandArray()
	fmt.Println(arr)
	mySort(arr, 0, 4)
	fmt.Println(arr)
}

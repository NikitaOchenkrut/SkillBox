package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func evenAndAdd(input [size]int) ([]int, []int) {
	var b []int
	var c []int

	for i := 0; i < size; i++ {
		if input[i]%2 != 0 {
			b = append(b, input[i])
		} else {
			c = append(c, input[i])
		}
	}
	return b, c
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var a [size]int

	for i := 0; i < size; i++ {
		a[i] = rand.Intn(10)
	}
	fmt.Println(evenAndAdd(a))
}

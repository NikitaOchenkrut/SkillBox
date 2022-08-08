package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func sortedInsert(input [n]int) [n]int {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if input[i] < input[j] {
				input[i], input[j] = input[j], input[i]
			}
		}
		fmt.Println(input)
	}
	return input
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var a [n]int

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10)
	}
	fmt.Println(a)
	fmt.Println(sortedInsert(a))

}

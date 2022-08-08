package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func buble(input [n]int) [n]int {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if input[j] < input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input

}

func main() {
	rand.Seed(time.Now().UnixNano())

	var b [n]int
	for i := 0; i < n; i++ {
		b[i] = rand.Intn(10)
	}
	fmt.Println(b)
	fmt.Println(buble(b))

}

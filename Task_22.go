package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 10

func main() {
	rand.Seed(time.Now().UnixNano())

	var a [n]int

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(10)
	}
	var num int
	fmt.Println("Какое число будем искать в массиве?: ")
	fmt.Scan(&num)
	fmt.Printf("Массив: %v\n", a)
	j := 0
	for i := 0; i < n-1; i++ {
		if a[i] == num {
			j = j + 1
		}
	}
	if j > 1 {
		if j < 5 {
			fmt.Printf("Число повторяется в массиве %v раза.", j)
		} else {
			fmt.Printf("Число повторяется в массиве %v раз.", j)
		}
	} else {
		fmt.Printf("Число повторяется в массиве %v раз.", j)
	}

}

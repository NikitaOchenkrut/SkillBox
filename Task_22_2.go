package main

import (
	"fmt"
	"math/rand"
	"time"
)

const n = 12

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
			j = i
			break
		}
	}
	fmt.Printf("Индекс первого числа по вашему запросу: %v ", j)

}

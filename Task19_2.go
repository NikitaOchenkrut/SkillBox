package main

import (
	"fmt"
)

const size = 10

func bubble(input [size]int) (newInput [size]int) {
	result := 0
	for i := 0; i < len(input); i++ {
		for i := 0; i < len(input); i++ {
			if input[i] == input[9] {
				break
			}
			if input[i] > input[i+1] {
				result = input[i]
				input[i] = input[i+1]
				input[i+1] = result
			}
		}
	}
	newInput = input
	return newInput
}
func main() {
	numbers := [size]int{1, 3, 6, 8, 34, 2, 54, 3, 65, 45}
	fmt.Println(bubble(numbers))

}

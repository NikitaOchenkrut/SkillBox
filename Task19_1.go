package main

import "fmt"

func generalArray(input1 [4]int, input2 [5]int) (newinput [9]int) {
	result := 0

	for i := 0; i < len(input1); i++ {
		newinput[i] = input1[i]
	}
	for i := 0; i < len(input2); i++ {
		newinput[i+len(input1)] = input2[i]
	}

	for i := 0; i < len(newinput); i++ {
		for i := 0; i < len(newinput); i++ {
			if newinput[i] == newinput[8] {
				break
			}
			if newinput[i] > newinput[i+1] {
				result = newinput[i]
				newinput[i] = newinput[i+1]
				newinput[i+1] = result
			}
		}
	}
	return newinput
}
func main() {
	number1 := [4]int{5, 7, 9, 6}
	number2 := [5]int{11, 10, 19, 10, 12}

	fmt.Println(number1, number2)
	fmt.Println(generalArray(number1, number2))
}

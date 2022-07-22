package main

import "fmt"

func determinant(input [3][3]int) int {
	matrix1 := [2][2]int{}
	matrix2 := [2][2]int{}
	matrix3 := [2][2]int{}
	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			matrix1[i-1][j-1] = input[i][j]
		}
	}
	for i := 1; i < 3; i++ {
		for j := 0; j < 2; j++ {
			matrix2[i-1][j] = input[i][j]
		}
	}
	for i := 1; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 0 {
				matrix3[i-1][j] = input[i][j]
			}
			if j == 2 {
				matrix3[i-1][j-1] = input[i][j]
			}
		}
	}
	fmt.Println(matrix1)
	fmt.Println(matrix2)
	fmt.Println(matrix3)
	matrix11 := [2]int{}
	matrix21 := [2]int{}
	matrix31 := [2]int{}
	res1 := 1
	res2 := 1
	result := 0

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == j {
				res1 = res1 * matrix1[i][j]
				matrix11[0] = res1
			} else {
				res2 = res2 * matrix1[i][j]
				matrix11[1] = -res2
			}
		}
	}
	res1 = input[0][0] * matrix11[0]
	res2 = input[0][0] * matrix11[1]
	result = res1 + res2 + result
	res1 = 1
	res2 = 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == j {
				res1 = res1 * matrix2[i][j]
				matrix21[0] = res1

			} else {
				res2 = res2 * matrix2[i][j]
				matrix21[1] = -res2

			}
		}
	}
	res1 = input[0][1] * matrix21[0]
	res2 = input[0][1] * matrix21[1]
	result = res1 + res2 + result
	res1 = 1
	res2 = 1
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if i == j {
				res1 = res1 * matrix3[i][j]
				matrix31[0] = res1
			} else {
				res2 = res2 * matrix3[i][j]
				matrix31[1] = -res2
			}
		}
	}
	res1 = input[0][2] * matrix31[0]
	res2 = input[0][2] * matrix31[1]
	result = res1 + res2 + result
	fmt.Println(matrix11)
	fmt.Println(matrix21)
	fmt.Println(matrix31)
	fmt.Println(result)

	return result
}

func main() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	determinant(matrix)

}

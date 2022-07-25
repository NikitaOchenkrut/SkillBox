package main

import "fmt"

const row = 3
const col = 3

func determinant(matrix [row][col]int) int {

	res1 := matrix[1][1]*matrix[2][2] - matrix[2][1]*matrix[1][2]
	res2 := matrix[1][0]*matrix[2][2] - matrix[2][0]*matrix[1][2]
	res3 := matrix[1][0]*matrix[2][1] - matrix[2][0]*matrix[1][1]

	determinant := matrix[0][0]*res1 - matrix[0][1]*res2 - matrix[0][2]*res3

	return determinant
}

func main() {
	matrix := [row][col]int{}
	fmt.Println("Введите данные матрицы: ")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%v строка, %v столбец.\n", i+1, j+1)
			fmt.Scan(&matrix[i][j])
		}
	}
	fmt.Println(determinant(matrix))

}

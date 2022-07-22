package main

import "fmt"

func multiplication(input1 [3][5]int, input2 [5][4]int) (newMatrix [3][4]int) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			for k := 0; k < 5; k++ {
				newMatrix[i][j] += input1[i][k] * input2[k][j]
			}

		}
	}
	fmt.Println(newMatrix)
	return newMatrix
}

func main() {
	matrixOne := [3][5]int{}
	matrixTwo := [5][4]int{}
	var res int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println("Введите в первую матрицу значение переменной X", i, j)
			fmt.Scan(&res)
			matrixOne[i][j] = res
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println("Введите во вторую матрицу значение переменной X", i, j)
			fmt.Scan(&res)
			matrixTwo[i][j] = res
		}
	}
	fmt.Println(matrixOne)
	fmt.Println(matrixTwo)
	multiplication(matrixOne, matrixTwo)
}

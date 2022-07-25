package main

import "fmt"

func resultS(x int16, y uint8, z float32) float32 {
	var S float32
	S = 2*float32(x) + float32(y*y) + 3/z
	return S

}
func main() {
	var a int16
	fmt.Println("Введите первую переменную: ")
	fmt.Scan(&a)
	var b uint8
	fmt.Println("Введите вторую переменную: ")
	fmt.Scan(&b)
	var c float32
	fmt.Println("Введите третью переменную: ")
	fmt.Scan(&c)
	fmt.Println(resultS(a, b, c))

}

package main

import "fmt"

const a = 15
const b = 10

func def(F func(int, int) int) int {
	defer F(a, b)
	fmt.Println("Вызываем анонимную функцию...")
	return F(a, b)
}
func main() {
	f1 := func(a, b int) int { return a * b }
	f2 := func(a, b int) int { return a / b }
	f3 := func(a, b int) int { return a + b }
	f1(a, b)
	f2(a, b)
	f3(a, b)
	fmt.Println(def(f1))
	fmt.Println(def(f2))
	fmt.Println(def(f3))

}

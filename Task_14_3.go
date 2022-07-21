package main

import "fmt"

func multiplication(x int)(a int){
  a = x * x
  return a
}
func plus(x int) (a int){
  a = x + x
  return a
}
func main(){
  fmt.Println("Введите число...")
  var a int
  fmt.Scan(&a)
  fmt.Println(multiplication(a))
  fmt.Println(plus(multiplication(a)))
}

package main

import "fmt"

func EvenOrOdd(x int) string{
  result := ""
  if x % 2 == 0{
    result = "True"
  }
  if x % 2 != 0{
     result = "False"
  }
  fmt.Println(result)
  return result
}

func main(){
  fmt.Println("Введите число.")
  var a int
  fmt.Scan(&a)
  EvenOrOdd(a)
}
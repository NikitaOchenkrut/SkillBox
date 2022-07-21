package main

import (
	"fmt"
)

func main() {
  var x int 

  fmt.Println("Давайте вычислим модуль из числа.")
  fmt.Println("Задайте число: ")
  fmt.Scan(&x)
  if x >= 0{
    fmt.Println("Модуль равен: " , x)
  }else if x < 0 {
    x = -x 
    fmt.Println("Модуль равен : " , x)
     
  }
  }
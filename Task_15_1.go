package main

import "fmt"

func main(){
  odd := 0
  even := 0
  var number [10]int
  for i:= 0 ; i < 10 ; i++{
    fmt.Println("Введите число: ")
    fmt.Scan(&number[i])
  }
  for i:= 0 ; i < 10 ; i++{
    if number[i] %2 == 0{
      even += 1
    }else{
      odd += 1
    }
}
  fmt.Printf("Четных чисел в массиве: %v\n", even)
  fmt.Printf("Нечетных чисел в массиве: %v\n", odd)
}
  
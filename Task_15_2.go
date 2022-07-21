package main

import "fmt"

func main(){
  var number [10]int
  fmt.Println("Введите в массив 10 чисел: ")
  for i:= 0 ; i < 10 ; i++{
    fmt.Println("Введите число: ")
    fmt.Scan(&number[i])
  }
  var newnumber [10]int
  j:= 9
  for i:= 0 ; i < 10 ; i++{
    newnumber[i] = number[j] 
    j -= 1
  }
  fmt.Println(newnumber)
}
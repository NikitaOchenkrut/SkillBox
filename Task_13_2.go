package main

import "fmt"

func swap(a, b *int){
  *a = *a + *b
  *b = *a - *b
  *a = *a - *b
}
func main(){
  fmt.Println("Введите два числа")
  var a int
  var b int
  fmt.Scan(&a)
  fmt.Scan(&b)
  fmt.Println(a, b)
  swap(&a, &b)
  fmt.Println(a, b)
  }

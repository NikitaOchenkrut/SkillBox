package main

import "fmt"

func summ(a, b int){
  sum := 0
  if a < b{
    for{
      if a % 2 == 0{
        sum += a
      }
      if a == b{
        fmt.Println("Насчитались")
        break
      }
      a += 1
    }
 }
  if b < a{
    for{
      if b % 2 == 0{
        sum += b 
      }
      if b == a{
        fmt.Println("Насчитались")
        break
      }
      b += 1
    }
  }
  fmt.Println(sum)
}
func main(){
  fmt.Println("Введите два числа, в диапазоне между ними мы возьмем все четные числа и суммируем.")
  var a int
  var b int
  fmt.Scan(&a)
  fmt.Scan(&b)
  summ(a, b)
  }
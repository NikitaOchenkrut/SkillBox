package main

import (
	"fmt"
)

func main() {
  fmt.Println("Задача три корзины.")

  fmt.Println("Сколько вмещает первая корзина?")
  var one int
  fmt.Scan(&one)
  fmt.Println("Сколько вмещает вторая корзина?")
  var two int
  fmt.Scan(&two)
  fmt.Println("Сколько вмещает третья корзина?")
  var three int
  fmt.Scan(&three)

  
  
  
  i:= 0
  for{
   
    i = i + 1 
    fmt.Println(i)
    if i == one {
      fmt.Println("Первая корзина заполнена.")
      continue
        if i == two {
        fmt.Println("Вторая корзина заполнена.")
        continue
          if i == three {
          fmt.Println("Третья корзина заполнена.")
          break
          }
         }
       } 
    }
  }
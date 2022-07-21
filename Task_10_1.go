package main

import (
	"fmt"
	"math"
)

func main() {
  fmt.Println("Разложение e^x в ряд Тейлора")
  fmt.Println("Введите x ")
  var x float64
  fmt.Scan(&x)
  fmt.Println("Введите точность знаков после запятой ")
  var n float64
  fmt.Scan(&n)
  epsilon := 1 / math.Pow(10, n) 
  
  var fakt float64
  fakt = 1.0
  result := 0.0
  oldResult := 0.0
  i := 1.0
  
  for {
    
    fakt = fakt * i
    result += math.Pow(x, i)/ fakt
    if (result - oldResult) < epsilon {
    fmt.Println(result+1)
    break
    }
    i++
    oldResult = result
  }
  }
  
  
  
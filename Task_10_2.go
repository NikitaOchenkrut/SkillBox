package main

import (
	"fmt"
  "math"
  
)

func main() {
  fmt.Println("Введите сумму зачисления в банк: ")
  var money float64
  fmt.Scan(&money)
  fmt.Println("Введите ежемесячный процент капитализации: ")
  var percent float64
  fmt.Scan(&percent)
  fmt.Println("Введите количество лет: ")
  var year float64
  fmt.Scan(&year)

  term := year * 12.0
  result := 0.0 
  partBank := 0.0
  i:= 1.0
  for {
    money = money + (money * (percent/100))
    result = money  
    money = math.Trunc(money * 100) / 100
    partBank += result-money
    if i == term {
      fmt.Printf("К концу срока ваш вклад составит %v банк заработал %v", money, partBank)
      break
    }
    i++
  }
  }
  
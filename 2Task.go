package main

import (
	"fmt"
)

func main() {

  var speed int
  var distance int

  fmt.Println("Из Москвы в Рязань выехал автомобиль. Нам нужно задать ему среднюю скорость, чтобы он успел доехать за 2 часа.")
  fmt.Println("")
  fmt.Println("Какая скорость у автомоболя? ")
  fmt.Scan(&speed)
  distance = 200 - (speed*2)

  
  if speed >= 100 {
    fmt.Println("Поздравляю, вы в Рязани")
  } else if speed<100{
    fmt.Println("Вы не успели доехать за 2 часа до Рязани, вам осталось преодолеть ", distance, "километра")
    
  }
  
  }
package main

import (
	"fmt"
)

func main() {
  var man int
  var barber int
  var work int
  var flaw int
  
  fmt.Println("Перед вами программа которая расчитает сколько нужно барберов для целого города.")
  fmt.Println("")
  fmt.Println("Сколько мужчин живет в городе? ")
  fmt.Scan(&man)
   fmt.Println("Сколько барберов работают в городе? ")
  fmt.Scan(&barber)
  
  work = barber*8*30
  
  flaw = (man - work)/(30*8)
  if flaw < 1{
    flaw++
  }
 
  if work>=man {
    fmt.Println ("Поздравляю, вашему городу хватает специалистов.", barber , "барберов, успеют постричь : ", work , "мужчин" )
  } else  {
    fmt.Println("Вам не хватает специалистов. Нужно нати ещё",flaw , "специалситов.")
    
  }
  
}
package main

import "fmt"

func main() {
  minFloat := 1
  maxFloat := 24
  cureentFloat := 1
  direction := 1

  maxPassengers := 2
  totalPassengers := 0

  passanger1 := 4
  passanger2 := 7
  passanger3 := 10

  for passanger1 != minFloat || passanger2 != minFloat || passanger3 != minFloat || cureentFloat != minFloat{
    cureentFloat += direction
    fmt.Println("Лифт находится на этаже:",cureentFloat,"Пассажиров в лифте:",totalPassengers)
    if direction == 1 && cureentFloat == maxFloat{
      direction = -1
    }else if cureentFloat == minFloat {
      direction = 1
      totalPassengers = 0
    }else if direction == -1{
      if passanger1 == cureentFloat && totalPassengers < maxPassengers {
        fmt.Println("Забираем первого пассажира с этажа",cureentFloat )
        passanger1 = 1
        totalPassengers++
        
      }
       if passanger2 == cureentFloat && totalPassengers < maxPassengers {
        fmt.Println("Забираем второг пассажира с этажа",cureentFloat )
         passanger2 = 1
         totalPassengers++
      }
       if passanger3 == cureentFloat && totalPassengers < maxPassengers {
        fmt.Println("Забираем третьего пассажира с этажа",cureentFloat )
         passanger3 = 1
         totalPassengers++
      }
      
      
    }
  } 
  fmt.Println("Все пасажиры на первом этаже.")
  }
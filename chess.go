package main

import (
	"fmt"
  "strings"
	"time"
)

func main() {
  
  fmt.Println("Сколько будет клеток в высоту?")
  var width int
  fmt.Scan(&width)
  
  fmt.Println("Сколько будет клеток в длинну?")
  var long int
  fmt.Scan(&long)
  for j := 0; j < width; j++ {
	 if j%2 != 0{
    fmt.Println(strings.Repeat("# ",long))
		time.Sleep(100 * time.Millisecond) 
   }else if j%2 == 0{
    fmt.Println(strings.Repeat(" #",long))
		time.Sleep(100 * time.Millisecond) 
   }  
  }
}
  
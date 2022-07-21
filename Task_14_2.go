package main

import (
	"fmt"
	"math/rand"
	"time"
)
func equation(x, y int) (int, int){
  x = 2 * x + 10
  y = -3 * y - 5
  fmt.Println(x, y)
  return x, y
}
func GetRandomPoints(n1, n2 int) (int, int){
  return rand.Intn(n1), rand.Intn(n2)
}
func main(){
  rand.Seed(time.Now().UnixNano())
  x1, y1 := GetRandomPoints(100, 100)
  equation(x1,y1)
  x2, y2 := GetRandomPoints(100, 100)
  equation(x2,y2)
  x3, y3 := GetRandomPoints(100, 100)
  equation(x3,y3)
  
}
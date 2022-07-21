package main

import "fmt"

var ONE_NUMBER = 8
var TWO_NUMBER = 16
var THREE_NUMBER = 36

func OneOpperation(x int)(a int){
  a = x + ONE_NUMBER
  return a
}
func TwoOpperation(x int)(a int){
  a = x + TWO_NUMBER
  return a
}
func ThreeOpperation(x int)(a int){
  a = x + THREE_NUMBER
  return a
}
func main(){
  fmt.Println("Введите число: ")
  var a int 
  fmt.Scan(&a)
  fmt.Println(OneOpperation(a))
  fmt.Println(TwoOpperation(OneOpperation(a)))
  fmt.Println(ThreeOpperation(TwoOpperation(OneOpperation(a))))
}

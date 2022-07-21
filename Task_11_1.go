package main

import (
	"fmt"
	"strings"
  
)

func main() {
  s := "Go is an Open source programming Language that makes it Easy to build simple, reliable, and efficient Software"
  fmt.Print("Определение количества слов, начинающихся с большой буквы в строке:\n", s, "\n")
  f := "QWERTYUIOPASDFGHJKLZXCVBNM"
  j := " "
  i := len(s)
  count := 0
  for {
    j = s[i:]
    if strings.Contains(f, j){
      count ++
    }
    s = s[:i]
    i=i-1
    if i == -1{
      break
    }
  }
  fmt.Println(count,"слов в строке начинаются с верхнего регистра.")
  
}
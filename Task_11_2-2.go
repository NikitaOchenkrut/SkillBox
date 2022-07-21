package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
  s := " a10 10 20b 20 30c30 30 dd "
  fmt.Println("Исходная строка : ", s)
   fmt.Println("Строка содержит числа в десятичном формате: ")
  for strings.Contains(s, " "){
    space := strings.Index(s, " ")
    i,err := strconv.ParseInt(s[:space], 10, 0)
      if err != nil {
      
      }else {
        fmt.Print(i, " ")
      }
    s = s[space + 1:]
  }
  
}
  
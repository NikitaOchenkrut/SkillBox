package main

import (
	"fmt"
  "math"
)

func main() {
  fmt.Println("Введите два числа.")
  var a,b int16
  fmt.Scan(&a)
  fmt.Scan(&b)
  
  fmt.Println("Выберем минимальный тип данных который нам подойдёт.")

  multiplication := int64(a)*int64(b)
  
  if multiplication > 0 {
    if multiplication < math.MaxUint8{
       fmt.Println(a, b, "результат uint8")
    }else if multiplication < math.MaxUint16{
       fmt.Println(a, b, "результат uint16")
    }else if multiplication < math.MaxUint32{
       fmt.Println(a, b, "результат uint32")
    }
  }else if multiplication < 0{
    if multiplication > math.MinInt8{
       fmt.Println(a, b, "результат int8")
    }else if multiplication > math.MinInt16{
       fmt.Println(a, b, "результат int16")
    }else if multiplication > math.MinInt32{
       fmt.Println(a, b, "результат int32")
    }  
  }
  }
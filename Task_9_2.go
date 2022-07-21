package main

import (
	"fmt"
)
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	
)

func main() {
  fmt.Println("Введите два числа.")
  var a,b int16
  fmt.Scan(&a)
  fmt.Scan(&b)
  
  fmt.Println("Выберем минимальный тип данных который нам подойдёт.")

  multiplication := int64(a*b)
  
  if multiplication > 0 {
    if multiplication < MaxUint8{
       fmt.Println(a, b, "результат uint8")
    }else if multiplication < MaxUint16{
       fmt.Println(a, b, "результат uint16")
    }else if multiplication < MaxUint32{
       fmt.Println(a, b, "результат uint32")
    }
  }else if multiplication < 0{
    if multiplication > MinInt8{
       fmt.Println(a, b, "результат int8")
    }else if multiplication > MinInt16{
       fmt.Println(a, b, "результат int16")
    }else if multiplication > MinInt32{
       fmt.Println(a, b, "результат int32")
    }  
  }
  }
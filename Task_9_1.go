package main

import (
	"fmt"  
)
const (
    
  MaxUint8 = 1<<8 - 1
  MaxUint16 = 1<<16 - 1
  MaxUint32 = 1<<32 - 1
)
func main() {
  fmt.Println("Посчитаем переполнения переменных.")
  for i := 0; i != MaxUint32; i++{
    
    if i == MaxUint8 {
      fmt.Println("Максимальное значение Uint8 %d ",i)
    }
    if i == MaxUint16 {
      fmt.Println("Максимальное значение Uint16 %d ",i)
    }
    if i == MaxUint32{
      break
    }
  }
  fmt.Println("Uint8 в диапазоне от 0 до Uint32 переполнен на ", MaxUint32-MaxUint8)
  fmt.Println("Uint8 в диапазоне от 0 до Uint32 переполнен на ", MaxUint32-MaxUint16)
  }

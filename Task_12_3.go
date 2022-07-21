package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  file, err := os.Create("12_3.txt")
  if err := os.Chmod("12_3.txt", 0444); err != nil{
    fmt.Println(err)
    return
  }
  writer := bufio.NewWriter(file)
  if err != nil{
   fmt.Println(err) 
  return
  }
  defer file.Close()
  writer.WriteString("Hello, World")
  }
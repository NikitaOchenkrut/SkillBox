package main

import (
	"fmt"
  "io"
  "os"
)

func main() {
  f, err := os.Open("Task12.txt")
  if err != nil {
    fmt.Println("Ошибка открытия файла", err) 
    return
  }
  defer f.Close()
  fi, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Size: %d\n", fi.Size())
	if fi.Size() < 1{
    fmt.Println("Файл пуст, никаких данных не содержит.")
  }

  buf := make([]byte, int(fi.Size()))
  if _, err := io.ReadFull(f, buf); err != nil{
    fmt.Println("Не смогли прочитать последовательность байтов из файла", err) 
    return
  }
  fmt.Printf("%s\n", buf) 
}
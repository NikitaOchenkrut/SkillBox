package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
  var b bytes.Buffer 
  
  i:=1
  for {
    fmt.Println("Сделайте запись")
    var word string
    fmt.Scan(&word)
    if word == "exit"{
      break
    }
    time_word := time.Now()
    
    fmt.Println("Запись номер: ", i,".", time_word.Format("02-Jan-2006. 15:04"), word)
    b.WriteString(fmt.Sprintf("Запись номер: ", i, time_word.Format("02-Jan-2006. 15:04"), word))
    b.WriteString(fmt.Sprintf("\n"))
    i += 1
    }
  fileName := "Task12.txt"
  if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil{
    panic(err)
  }
  file, err := os.Open(fileName)
  if err!= nil{
    panic(err)
  }
  defer file.Close()
   fi, err := file.Stat()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Size: %d\n", fi.Size())
	if fi.Size() < 1{
    fmt.Println("Файл пуст, никаких данных не содержит.")
  }
  result, err := ioutil.ReadAll(file)
  if err != nil{
    panic(err)
  }
  fmt.Println(string(result))
}

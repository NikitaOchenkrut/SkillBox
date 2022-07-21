package main

import (
	"fmt"
  "time"
  "os"
 
  )

func main() {
  file, err := os.Create("Task12.txt")
  if err != nil{
    fmt.Println(err)
    return
  } 
  defer file.Close()
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
    file.WriteString(fmt.Sprintf("Запись номер: ", i, time_word.Format("02-Jan-2006. 15:04"), word))
    file.WriteString(fmt.Sprintf("\n"))
    i += 1
    }
}

  
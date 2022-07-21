package main

import (
	"fmt"
  "time"
  "os"
  "strconv"
 
  )

func main() {
  file, err := os.Create("Task12_1.txt")
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
    var n string
    n = strconv.Itoa(i)
    fmt.Println("Запись номер: ", i,".", time_word.Format("02-Jan-2006. 15:04"), word)
    file.WriteString(" Запись номер ")
    file.WriteString(n)
    file.WriteString(".")
    file.WriteString(time_word.Format("02-Jan-2006. 15:04"))
    file.WriteString(" ")
    file.WriteString(word)
    file.WriteString("\n")
    i += 1
    }
}

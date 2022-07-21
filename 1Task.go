package main

import (
	"fmt"
)

func main() {

  var price int 
  var sum int

  fmt.Println("Укажите сумму первого товара: ")
  fmt.Scan(&price)
  sum = price
  fmt.Println("Укажите сумму второго товара: ")
  fmt.Scan(&price)
  sum = sum + price
  fmt.Println("Укажите сумму третьего товара: ")
  fmt.Scan(&price)
  sum = sum + price

  if sum > 20000 {
    sum = sum/100*80
    fmt.Println("Сумма товаров превышает 20 000, вы получаете скидку 20% ")
  } else if  sum > 10000 {
    sum = sum/100*90
    fmt.Println("Сумма товаров превышает 10 000, вы получаете скидку 10% ")
    
  }
  
  fmt.Println("------------")
  fmt.Println("Итого к оплате:", sum)
  }
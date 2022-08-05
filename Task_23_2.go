package main

import (
	"fmt"
	"strings"
)

func parseTest(sentences [4]string, chars [5]rune) string {
	result := " "

	for i := 0; i < 4; i++ {
		for j := 0; j < 5; j++ {
			if strings.Contains(sentences[i], string(chars[j])) == true {
				fmt.Println(string(chars[j]), " Входит в строку: ", sentences[i])
				result = sentences[i]
				for i := 0; i < len(result); i++ {
					if string(result[i]) == string(chars[j]) {
						fmt.Println("Позиция: ", i)
						break

					}
				}

			}
		}
	}
	return " "
}

func main() {
	my_sentences := [4]string{"HELLO", "Bye-Bye", "Ma name is", "HELLO WORLD"}
	my_chars := [5]rune{'H', 'E', 'L', 'm', 'n'}

	fmt.Println(parseTest(my_sentences, my_chars))
}

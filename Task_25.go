package main

import (
	"flag"
	"fmt"
)

func main() {
	var str string
	var substr string

	flag.StringVar(&str, "str", "строка для поиска", "string to search in")
	flag.StringVar(&substr, "substr", "поиска", "string to search for")
	flag.Parse()

	fmt.Println(isSubstring(&substr, &str))
	result := isSubstring(&substr, &str)

	if result == true {
		fmt.Printf("Строка: %v , включает в себя подстроку: %v ", str, substr)
	}

}

func isSubstring(substr *string, str *string) bool {
	substrRunes := []rune(*substr)
	strRunes := []rune(*str)

	j := 0
	for _, r := range strRunes {
		if r == substrRunes[j] {
			j++
			if j == len(substrRunes) {
				return true
			}
		} else {
			j = 0
		}
	}
	return false
}

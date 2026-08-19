package vowels

import "fmt"

func Vowels() string {
	var ch string
	fmt.Println("enter a character")
	fmt.Scan(&ch)
	if ch == "a" || ch == "e" || ch == "i" || ch == "o" || ch == "u" {
		fmt.Println("it is a vowel")
	} else {
		fmt.Println("it is a constant")
	}
	return ch
}

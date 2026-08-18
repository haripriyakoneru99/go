package main

import "fmt"

func main() {
	number := [5]int{1, 5, 8, 9, 6}
	fmt.Println(number)
	fmt.Println(number[1])
	fmt.Println(number[2])
	fmt.Println(number[3])
	number[2] = 10
	fmt.Println(number[2])
	fmt.Println(len(number))
}

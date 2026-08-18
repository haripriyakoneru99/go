package array

import "fmt"

func Arr() [10]int {
	var number [10]int
	for i := 0; i < 10; i++ {
		fmt.Println("enter number")
		fmt.Scan(&number[i])
	}
	return number
}

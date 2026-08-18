package array

import "fmt"

func Ar() [5]int {

	var number [5]int

	for i := 0; i < 5; i++ {

		fmt.Println("Enter number")

		fmt.Scan(&number[i])
	}
	fmt.Println(number[0])

	return number
}

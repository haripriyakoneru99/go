package swap

import "fmt"

func Sw() (int, int) {
	var a int
	var b int
	var temp int
	fmt.Println("enter a value")
	fmt.Scan(&a)
	fmt.Println("enter b value")
	fmt.Scan(&b)
	temp = a
	a = b
	b = temp
	return a, b
}

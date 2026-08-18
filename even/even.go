package even

import "fmt"

func Ev() int {
	var n int
	fmt.Println("enter a value")
	fmt.Scan(&n)
	if n%2 == 0 {
		fmt.Println("it is even")
	} else {
		fmt.Println("it is zero/odd")
	}
	return n
}

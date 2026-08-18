package main

import (
	"GOLANG/array"
	"GOLANG/even"
	"GOLANG/pointers"
	"fmt"
)

func main() {
	var n int
	fmt.Println("enter n value")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
	fmt.Println(pointers.Point())
	fmt.Println(array.Ar())
	fmt.Println(array.Arr())
	fmt.Println(even.Ev())
}

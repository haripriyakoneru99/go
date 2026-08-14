package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter n value")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Println(i)
	}
}

package pointers

import "fmt"

func Point() int {
	a := 1
	p := &a
	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)
	return *p
}

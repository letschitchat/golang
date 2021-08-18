package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:4]
	fmt.Println(b)

	c := []int{6, 7, 8}
	fmt.Println(c)

	darr := [...]int{50, 61, 83, 100, 68, 53, 70}
	dslice := darr[2:5]
	fmt.Println("dslice is ", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("after is ", darr)
}

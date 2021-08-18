package main

import "fmt"

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside func", num)
}

func main() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a[0])

	b := [3]int{4, 5, 6}
	fmt.Println(b)

	c := [...]string{"james", "koby", "malong"}
	d := c
	d[0] = "zhangjk"
	fmt.Println("c is ", c)
	fmt.Println("d is ", d)

	num := [...]int{5, 6, 7, 8, 9}
	fmt.Println("before func", num)
	changeLocal(num)
	fmt.Println("after cun,length", num, len(num))

}

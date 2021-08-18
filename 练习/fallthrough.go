package main

import "fmt"

func number() int {
	num := 15 * 3
	return num
}

func main() {
	switch num := number(); {
	case num < 50:
		fmt.Println("what are u")
		fallthrough
	case num < 100:
		fmt.Println("hello")
	}
}

package main

import "fmt"

func main() {
	a := [...]float64{32.21, 42.21, 52.21}
	sum := float64(0)
	for i, v := range a {
		fmt.Println(i, v, sum)
		sum += v
	}
	fmt.Println(sum)
}

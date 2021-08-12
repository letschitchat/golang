package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Printf(" %d", i)
	}

	for k, v := 10, 1; k <= 19 && v <= 10; k, v = k+1, v+1 {
		fmt.Printf("\n %d * %d = %d", k, v, k*v)
	}

}

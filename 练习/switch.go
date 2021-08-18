package main

import "fmt"

func main() {
	letter := "o"
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Printf("zk")
	default:
		fmt.Println("nozk")
	}

	num := 15
	switch {
	case num > 10 && num < 20:
		fmt.Println("\nabcdefg")

	}
}

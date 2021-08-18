package main

import "fmt"

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
		fmt.Println(v1)
	}
}

func main() {
	a := [3][2]string{
		{"x", "y"},
		{"g", "h"},
		{"u", "k"},
	}
	printarray(a)

	var b [3][2]string
	b[0][0] = "apple"
	b[0][1] = "samsung"
	b[1][0] = "microsoft"
	b[1][1] = "google"
	b[2][0] = "AT&T"
	b[2][1] = "T-Mobile"
	fmt.Printf("\n")
	printarray(b)

}

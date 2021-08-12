package main

import "fmt"

func main() {
	c1 := complex(3, 8)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum: ", cadd)
	cmul := c1 * c2
	fmt.Println("product: ", cmul)

	k := 10
	var j float64 = float64(k)
	fmt.Println(j)

}

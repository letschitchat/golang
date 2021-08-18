package main

import "fmt"

func main() {
	i := make([]int, 5, 5)
	fmt.Println(i, len(i), cap(i))

	cars := []string{"honda", "ford"}
	fmt.Println("cars: ", cars)
	cars = append(cars, "aodi")
	fmt.Println("cars append :", cars)

	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "abcd", "efgh")
		fmt.Println("names content ", names)
	}

	pls := [][]string{
		{"c", "c++"},
		{"go", "rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

}

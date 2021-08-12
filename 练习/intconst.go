package main

import (
	"fmt"
)

func calcBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func main() {
	const a = 5
	var intVar int = a
	var int32Var int32 = a
	var float64Var float64 = a
	var complex64Var complex64 = a
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	price, no := 30, 100
	totalPrice := calcBill(price, no)
	fmt.Println(totalPrice)

	length, width := 10.5, 20.5
	area, perimeter := rectProps(length, width)
	fmt.Println(area, perimeter)
}

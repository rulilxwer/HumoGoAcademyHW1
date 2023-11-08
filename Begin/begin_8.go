package main

import "fmt"

func main() {
	var a float64
	var b float64
	fmt.Println("Введите число a :")
	fmt.Scan(&a)
	fmt.Println("Введите число b :")
	fmt.Scan(&b)
	var SA float64 = (a + b) / 2
	fmt.Println("Среднее арифметическое чисел", a, b, " равен", SA)
}

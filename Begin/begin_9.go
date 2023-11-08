package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64
	fmt.Println("Введите число a :")
	fmt.Scan(&a)
	fmt.Println("Введите число b :")
	fmt.Scan(&b)
	var SG float64 = math.Sqrt(a * b)
	fmt.Println("Среднее геометрическое чисел", a, b, " равен", SG)
}

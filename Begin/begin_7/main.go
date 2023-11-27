package main

import (
	"fmt"
	"math"
)

func main() {
	var R float64
	fmt.Println("Введите радиус :")
	fmt.Scan(&R)
	var L float64 = math.Round(2 * 3.14 * R)
	var S float64 = math.Round(3.14 * R * R)
	fmt.Println("Длина окружности ", R, " равен", L, " а его площадь равна", S)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var L float64
	fmt.Println("Введите расстояние в саниметрах:")
	fmt.Scan(&L)
	var M float64 = math.Round(L / 100)
	fmt.Println("Расстояние в сантиметрах", L, ", в метрах", M)
}

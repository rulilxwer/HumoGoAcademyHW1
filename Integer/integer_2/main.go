package main

import (
	"fmt"
	"math"
)

func main() {
	var M float64
	fmt.Println("Введите массу в граммах ")
	fmt.Scan(&M)
	var T float64 = math.Round(M / 1000)
	fmt.Println("Масса в граммах", M, " в килограммах", T)
}

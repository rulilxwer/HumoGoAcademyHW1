package main

import (
	"fmt"
	"math"
)

func main() {
	var B float64
	fmt.Println("Введите размер в байтах:")
	fmt.Scan(&B)
	var K float64 = math.Round(B / 1024)
	fmt.Println("Размер в байтах", B, "  в килобайтах", K)
}

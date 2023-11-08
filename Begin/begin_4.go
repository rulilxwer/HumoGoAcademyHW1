package main

import "fmt"

func main() {
	var d float32
	fmt.Println("Введите длину окружности")
	fmt.Scan(&d)
	var dlina float32 = (d) * 3.14
	fmt.Println("Длина окружности", d, " равен", dlina)
}

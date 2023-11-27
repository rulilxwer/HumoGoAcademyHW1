package main

import "fmt"

func main() {
	var A int
	var B int
	fmt.Print("Задаем значения А и В")
	fmt.Scan(&A)
	fmt.Scan(&B)
	var remainder int = A % B

	fmt.Println("Длина незанятой части отрезка A равна", remainder)
}

package main

import "fmt"

func main() {
	var A int
	var B int
	fmt.Print("Задаем значения А и В")
	fmt.Scan(&A)
	fmt.Scan(&B)
	remainder := A / B

	fmt.Println(remainder)
}

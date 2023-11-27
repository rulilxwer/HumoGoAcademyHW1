package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите значение а ")
	fmt.Scan(&a)
	var b int
	fmt.Println("Введите значение b ")
	fmt.Scan(&b)

	var result bool

	result = a >= 0 || b < -2

	fmt.Println(result)
}

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

	result = a%2 == 1 || b%2 == 1

	fmt.Println(result)
}

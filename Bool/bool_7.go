package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите значение а ")
	fmt.Scan(&a)

	var b int
	fmt.Println("Введите значение b ")
	fmt.Scan(&b)

	var c int
	fmt.Println("Введите значение c ")
	fmt.Scan(&c)

	var result bool

	result = a < b && b < c || a > b && b > c

	fmt.Println(result)
}

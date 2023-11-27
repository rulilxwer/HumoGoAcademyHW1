package main

import "fmt"

func main() {
	var a int
	var result bool
	fmt.Println("Введите число : ")
	fmt.Scan(&a)
	result = a%2 == 1
	fmt.Println(result)
}

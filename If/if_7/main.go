package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите значение a")
	fmt.Scan(&a)
	var b int
	fmt.Println("Введите значение b")
	fmt.Scan(&b)
	if a > b {
		fmt.Println("№2", b)
	} else {
		fmt.Println("№1", a)
	}
}

package main

import "fmt"

func main() {
	var a float64
	fmt.Println("Введите значение a")
	fmt.Scan(&a)
	var b float64
	fmt.Println("Введите значение b")
	fmt.Scan(&b)
	if a > b {
		fmt.Println("Значение a равно :", b, "Значение b равно :", a)
	} else {
		fmt.Println("Значение a равно :", a, "Значение b равно :", b)
	}
}

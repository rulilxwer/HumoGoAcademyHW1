package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите значение a")
	fmt.Scan(&a)
	var b int
	fmt.Println("Введите значение b")
	fmt.Scan(&b)
	var c int
	fmt.Println("Введите значение c")
	fmt.Scan(&c)
	if a > 0 && b > 0 && c > 0 {
		fmt.Println("Количество положительных чисел равно 3")
	} else if (a < 0 && b > 0 && c > 0) || (a > 0 && b < 0 && c > 0) || (a > 0 && b > 0 && c < 0) {
		fmt.Println("Количество положительных чисел равно :", 2, ", а количество отрицательных равен :", 1)
	} else if a < 0 && b < 0 && c < 0 {
		fmt.Println(" количество отрицательных равен 3 ")
	} else {
		fmt.Println("Количество положительных чисел равно :", 1, ", а количество отрицательных равен :", 2)
	}

}

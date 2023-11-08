package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Println("Введите сторону ребра a :")
	fmt.Scan(&a)
	fmt.Println("Введите длину ребра b :")
	fmt.Scan(&b)
	fmt.Println("Введите длину ребра c :")
	fmt.Scan(&c)
	var obyom = a * b * c
	var ploshad = 2 * (a*b + b*c + a*c)
	fmt.Println("Объём куба", a, b, c, " равен", obyom, " а его площадь равна", ploshad)
}

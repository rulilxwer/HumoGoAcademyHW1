package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите длину ребра куба :")
	fmt.Scan(&a)
	var obyom = a * a * a
	var ploshad = 6 * (a * a)
	fmt.Println("Объём куба", a, " равен", obyom, " а его площадь равен", ploshad)
}

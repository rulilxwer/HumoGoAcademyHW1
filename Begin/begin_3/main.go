package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Введите сторону a :")
	fmt.Scan(&a)
	fmt.Println("Введите сторону b :")
	fmt.Scan(&b)
	var ploshad = a * b
	var perimetr = 2 * (a + b)
	fmt.Println("Площадь прямоугольника со сторонами", a, "и", b, " равен", ploshad, " а", " периметр равен", perimetr)
}

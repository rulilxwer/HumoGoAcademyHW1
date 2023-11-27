package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите сторону квадрата:")
	fmt.Scan(&a)
	var ploshad = a * a
	fmt.Println("Площадь квадрата со стороной", a, " равен", ploshad)
}

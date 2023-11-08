package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите сторону квадрата:")
	fmt.Scan(&a)
	var perimeter = 4 * a
	fmt.Println("Периметр квадрата со стороной", a, " равен", perimeter)
}

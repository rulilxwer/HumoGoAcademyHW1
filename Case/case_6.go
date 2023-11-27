package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите длину отрезка")
	fmt.Scan(&a)
	var o int
	fmt.Println("Выберите единицу измерения длины (1-5)")
	fmt.Scan(&o)
	unitOfLength := o
	switch unitOfLength {
	case 1:
		fmt.Println(a, " дециметров равно", a/10, " метров")
	case 2:
		fmt.Println(a, " киломеров равно ", a*1000, " метров")
	case 3:
		fmt.Println(a, " метров")
	case 4:
		fmt.Println(a, "миллиметров равно", a/1000, " метров")
	case 5:
		fmt.Println(a, "сантиметров", a/100, " метров")
	default:
		fmt.Println("Вветиде число в диапазоне от 1 до 5")
	}
}

package main

import "fmt"

func main() {
	var month int
	fmt.Println("Введите число соотвествующее месяцу ")
	fmt.Scan(&month)
	monthOfTheYear := month
	switch monthOfTheYear {
	case 1:
		fmt.Println("Январь , зима")
	case 2:
		fmt.Println("Февраль , зима")
	case 3:
		fmt.Println("Март , весна")
	case 4:
		fmt.Println("Апрель , весна")
	case 5:
		fmt.Println("Май , весна")
	case 6:
		fmt.Println("Июнь , лето")
	case 7:
		fmt.Println("Июль , лето")
	case 8:
		fmt.Println("Август , лето")
	case 9:
		fmt.Println("Сентябрь , осень")
	case 10:
		fmt.Println("Октябрь , осень")
	case 11:
		fmt.Println("Ноябрь , осень")
	case 12:
		fmt.Println("Декабрь , зима")
	default:
		fmt.Println("Такого месяца не существует")
	}
}

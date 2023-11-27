package main

import "fmt"

func main() {
	var month int
	fmt.Println("Введите число соотвествующее месяцу ")
	fmt.Scan(&month)
	monthOfTheYear := month
	switch monthOfTheYear {
	case 1:
		fmt.Println("Январь , 31")
	case 2:
		fmt.Println("Февраль , 28")
	case 3:
		fmt.Println("Март , 31")
	case 4:
		fmt.Println("Апрель , 30")
	case 5:
		fmt.Println("Май , 31")
	case 6:
		fmt.Println("Июнь , 30")
	case 7:
		fmt.Println("Июль , 31")
	case 8:
		fmt.Println("Август , 31")
	case 9:
		fmt.Println("Сентябрь , 30")
	case 10:
		fmt.Println("Октябрь , 31")
	case 11:
		fmt.Println("Ноябрь , 30")
	case 12:
		fmt.Println("Декабрь , 31")
	default:
		fmt.Println("Такого месяца не существует")
	}
}

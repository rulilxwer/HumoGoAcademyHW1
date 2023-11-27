package main

import "fmt"

func main() {
	var A float64
	fmt.Println("Введите число A ")
	fmt.Scan(&A)
	var B float64
	fmt.Println("Введите число B ")
	fmt.Scan(&B)
	var n int
	fmt.Println("Выберите номер действия ")
	fmt.Scan(&n)
	if B == 0 {
		fmt.Println("Значение B не может быть нулевым")
	} else {
		arithmeticOperation := n
		switch arithmeticOperation {
		case 1:
			fmt.Println("Сложение", A+B)
		case 2:
			fmt.Println("Вычитание", A-B)
		case 3:
			fmt.Println("Умножение", A*B)
		case 4:
			fmt.Println("Деление", A/B)
		default:
			fmt.Println("Выберите номер действие от 1 до 4 !")
		}
	}
}

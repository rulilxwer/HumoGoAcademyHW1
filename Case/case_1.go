package main

import "fmt"

func main() {
	var day int
	fmt.Println("Введите число соотвествующее дню недели")
	fmt.Scan(&day)
	dayOfWeek := day
	switch dayOfWeek {
	case 1:
		fmt.Println("Понедельник")
	case 2:
		fmt.Println("Вторник")
	case 3:
		fmt.Println("Среда")
	case 4:
		fmt.Println("Четверг")
	case 5:
		fmt.Println("Пятница")
	case 6:
		fmt.Println("Суббота")
	case 7:
		fmt.Println("Воскресенье")
	default:
		fmt.Println("Выберите число от 1 до 7")
	}

}

package main

import "fmt"

func main() {
	var K int
	fmt.Println("Введите оценку")
	fmt.Scan(&K)
	mark := K
	switch mark {
	case 1:
		fmt.Println("плохо")
	case 2:
		fmt.Println("неудовлетворительно")
	case 3:
		fmt.Println("удовлетворительно")
	case 4:
		fmt.Println("хорошо")
	case 5:
		fmt.Println("отлично")
	default:
		fmt.Println("такой оценки не существует")
	}
}

package main

import "fmt"

func checkSign( array:=make([]int)) string {
	for i := 0; i < len(array)-1; i++ {
		if array[i]*array[i+1] > 0 {
			return "Yes"
		}
	}
	return "No"
}

func main() {
	var n int
	fmt.Println("Введите элементы массива через пробел:")
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	result := checkSign(array)
	fmt.Println(result)

}

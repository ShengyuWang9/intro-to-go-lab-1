package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
}

func mapSlice(f func(a int) int, slice []int) {

}

func mapArray(f func(a int) int, array [3]int) {

}

func main() {
	////question 3a
	//var array = [3]int{5, 10, 15}
	//for _, num := range array {
	//	fmt.Println(addOne(num))
	//}

	//question 3b
	intsSlice := []int{1, 2, 3}
	for _, num1 := range intsSlice {
		fmt.Println(addOne(num1))
	}

	intsArray := [3]int{1, 2, 3}
	for _, num2 := range intsArray {
		fmt.Println(addOne(num2))
	}
}

package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) []int {
	for _, num := range slice {
		slice = append(slice, num)
	}
	return slice
}

func mapSlice(f func(a int) int, slice []int) {
	for i, num := range slice {
		slice[i] = f(num)
	}
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	for i, num := range array {
		array[i] = f(num)
	}
	return array
}

func main() {
	////question 3a
	//var array = [3]int{5, 10, 15}
	//for _, num := range array {
	//	fmt.Println(addOne(num))
	//}

	//question 3b
	//doesn't follow the request
	//intsSlice := []int{1, 2, 3}
	//for _, num1 := range intsSlice {
	//	fmt.Println(addOne(num1))
	//}
	//
	//intsArray := [3]int{1, 2, 3}
	//for _, num2 := range intsArray {
	//	fmt.Println(addOne(num2))
	//}

	//intsSlice := []int{1, 2, 3}
	//mapSlice(addOne, intsSlice)
	//fmt.Println(intsSlice)
	//intsArray := [3]int{1, 2, 3}
	//intsArray = mapArray(addOne, intsArray)
	//fmt.Println(intsArray)

	//question 3c
	//intsSlice := []int{1, 2, 3, 4, 5}
	//mapSlice(addOne, intsSlice)
	//fmt.Println(intsSlice)
	//intsArray := [5]int{1, 2, 3, 4, 5}
	//intsArray = mapArray(addOne, intsArray)
	//fmt.Println(intsArray)

	//the outcome shows that array in function can't accept the different requirement, but slices can
	//when modify the input parameter requirement, the program can run

	////question 3d
	//var intsSlice = []int{2, 3, 4, 5, 6}
	//newSlice := intsSlice[1:3]
	//fmt.Println(newSlice)
	//for _, num := range newSlice {
	//	fmt.Println(square(num))
	//}

	//question 3e
	intsSlice := []int{5, 6, 7}
	fmt.Println(double(intsSlice))

}

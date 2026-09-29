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
	var array = [3]int{5, 10, 15}
	for _, num := range array {
		go addOne(num)
		fmt.Println(num)
	}

}

package main

import "fmt"

func main() {
	num := []int{2, 3, 4}
	sum := 0
	for num := range num {
		sum += num
	}
	fmt.Println("Sum:", sum)

}

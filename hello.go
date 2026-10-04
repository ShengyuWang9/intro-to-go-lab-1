package main

import (
	"fmt"
	"time"
)

//func main() {
//	for i := 0; i < 20; i++ {
//		fmt.Println("Hello World")
//	}
//}

func sayHelloWorld(i int) {
	fmt.Println("Hello from goroutine", i)
}

func main() {
	for i := 0; i <= 4; i++ {
		go sayHelloWorld(i)

	}
	time.Sleep(time.Second)

}

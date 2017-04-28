package main

import "fmt"

func main() {
	value := make(chan int, 5)
	go produce(10, value)
	consume(value)
}

func produce(num int, value chan int) {
	defer close(value)
	defer fmt.Println("done!")

	for i := 0; i < num; i++ {
		value <- i
	}
}

func consume(value chan int) {
	for i := range value {
		fmt.Println(i)
	}
}

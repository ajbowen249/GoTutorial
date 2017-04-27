package main

import "fmt"

func main() {
	fmt.Println("hello, world")
	var x int
	x = 3

	y := 1.23

	z := false

	fmt.Printf("x: %v y: %v z: %v\n", x, y, z)

	var hiddenInt interface{}
	hiddenInt = 12
	exposedInt := hiddenInt.(int)

	floatyInt := 15.0
	truncatedInt := int(floatyInt)

	fmt.Printf("exposedInt: %v truncatedInt: %v\n", exposedInt, truncatedInt)
}

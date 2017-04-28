package main

import "fmt"

func main() {
	doCompare(12, 1, greater, printIftrue)

	printIfFalse := func(value bool) {
		if !value {
			fmt.Println("false")
		} else {
			fmt.Println("not false")
		}
	}

	doCompare(12, 1, greater, printIfFalse)

	doCompare(12, 12, func(x int, y int) bool {
		return x == y
	}, func(equal bool) {
		if equal {
			fmt.Println("equal")
		} else {
			fmt.Println("not equal")
		}
	})
}

func greater(x int, y int) bool {
	return x > y
}

func printIftrue(greater bool) {
	if greater {
		fmt.Println("true")
	} else {
		fmt.Println("not true")
	}
}

func doCompare(x int, y int, compare func(int, int) bool, act func(bool)) {
	act(compare(x, y))
}

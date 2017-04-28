package main

import "fmt"

func main() {
	val := 3
	switch {
	case val == 1:
		fmt.Println("val is 1")
	case val == 3:
		fmt.Println("val is 3")
	case val > 2:
		fmt.Println("val is more than 2")
		fallthrough
	case val > 1:
		fmt.Println("val is more than 1")
		fallthrough
	case val > 0:
		fmt.Println("val is more than 0")
	case val == 9:
		fmt.Println("val is 9")
	default:
		fmt.Println("val is something")
	}
}

func max2(val1 int, val2 int) (int, bool) {
	if val1 == val2 {
		return val1, true
	} else if val1 > val2 {
		return val1, false
	} else {
		return val2, false
	}
}

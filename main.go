package main

import "fmt"

func main() {
	bigger, _ := max2(13, 13)
	fmt.Printf("%v is bigger.\n", bigger)
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

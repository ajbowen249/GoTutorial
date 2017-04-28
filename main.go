package main

import "fmt"

func main() {
	point1 := point{y: 6}
	point2 := point{}
	Print(point1)

	point2.TranslateX(5)
	Print(point2)

	point1.TranslateY(-3)
	Print(point1)
}

type point struct {
	x int
	y int
}

func (p point) ToString() string {
	return fmt.Sprintf("(%v, %v)", p.x, p.y)
}

func (p point) TranslateX(delta int) {
	p.x += delta
}

func (p point) TranslateY(delta int) {
	p.y += delta
}

func Print(val interface {
	ToString() string
}) {
	fmt.Println(val.ToString())
}

package main

import "fmt"

type shape interface {
	area() float64
}

type triangle struct{
	height float64
	base float64
}

type square struct{
	sideLength float64
}

func (t triangle) area() float64 {
	return (t.base * t.height) / 2
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.area())
}

func main() {
	s := square{sideLength: 3}
	t := triangle{base: 2, height: 2}
	printArea(s)
	printArea(t)
}

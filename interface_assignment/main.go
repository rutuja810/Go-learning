package main

import "fmt"
 
type shape interface {
	getArea() float64
}

type rectangle struct {
	rlength float64
	rbreadth float64

}

type square struct {
	slength float64

}

func (rect rectangle) getArea() float64 {
	return rect.rlength * rect.rbreadth
}

func (sq square) getArea() float64 {
	return sq.slength * sq.slength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
func main() {
	s := square{20}
	r := rectangle{12, 10}
	printArea(s)
	printArea(r)
}
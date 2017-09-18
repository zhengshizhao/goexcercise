package main

import (
	"fmt"
	"math"
)

type square struct {
	length float32
	width  float32
}
type circle struct {
	radius float32
}
type shape interface {
	area() float32
}

func (s square) area() float32 {
	return s.length * s.width
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}
func info(sh shape) {
	fmt.Println(sh.area())
}
func main() {
	square1 := square{3.5, 5}
	circle1 := circle{7.8}
	// var is_rount bool
	i := 1
	j := 4.6
	is_rount := true
	fmt.Printf("General:")
	fmt.Printf("%%v i, j: %v, %v\n", i, j)
	fmt.Printf("%%[2]v [1]v: %[2]v, %[1]v\n", i, j)
	// fmt.Println(fmt.Sprint("%[3]*.[2]*[1]f ", 12.0, 2, 6))
	fmt.Printf("%[3]*.[2]*[1]f\n", 12.0, 2, 6)
	fmt.Printf("%%v: %v\n", square1)
	fmt.Printf("%%T: %T\n", square1)
	fmt.Printf("%%#v: %#v\n", square1)
	fmt.Printf("%%+v: %+v\n", square1)
	fmt.Printf("%%: %%\n")
	fmt.Printf("%%t %t\n", is_rount)



	info(square1)
	info(circle1)
}

// HANDS ON 1
// create a type square
// create a type circle
// attach a method to each that calculates area and returns it
// create a type shape which defines an interface as anything which has the area method
// create a func info which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

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
	info(square1)
	info(circle1)
}

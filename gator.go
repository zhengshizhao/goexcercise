/*
1) Create a new type called “gator”. The underlying type of “gator” is an int.
Using var, declare an identifier “g1” as type gator (var g1 gator).
Assign a value to “g1”. Print out “g1”. Print the type of “g1” using fmt.Printf(“%T\n”, g1)

2)  Using var, declare an identifier “x” as type int (var x int).
Print out “x”. Print the type of “x” using fmt.Printf(“%T\n”, x)

3) assince x to g1 using CONVERSION

4)Now add a method to type gator with this signature ...
	greeting()
… and have it print “Hello, I am a gator”. Create a value of type gator. Call the greeting() method from that value.

*/
package main

import "fmt"

type gator int

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func main() {
	var g1 gator
	var x int
	x = 4
	g1 = gator(x)
	fmt.Println("x value:", x)
	fmt.Printf("type of x is %T\n", x)

	fmt.Println("g1 value:", g1)
	g1.greeting()
	fmt.Printf("g1 type is %T\n", g1)

}

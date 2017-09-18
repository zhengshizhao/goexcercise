/*Create a new type called “transportation”.
The underlying type of this new type is interface. An interface defines functionality.
Said another way, an interface defines behavior.
For this interface, any other type that has a method with this signature …
	transportationDevice() string
… will automatically, implicitly implement this interface.
Create a func called “report” that takes a value of type “transportation” as an argument.
The func should print the string returned by “transportationDevice()” being called for any type
that implements the “transportation” interface.
Call “report” passing in a value of type truck.
Call “report” passing in a value of type sedan.*/

package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func (t truck) transportationDevice() string {
	return fmt.Sprintln("this is a", t.vehicle.color, "truck")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("this is a", s.vehicle.color, "sedan")
}
func main() {
	myTruck := truck{vehicle{4, "blue"}, true}
	mySedan := sedan{vehicle{2, "red"}, false}

	report(myTruck)
	report(mySedan)

}

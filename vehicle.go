/*
1) Create a new type: vehicle. The underlying type is a struct. 
The fields: doors, color. Create two new types: truck & sedan. 
The underlying type of each of these new types is a struct. 
Embed the “vehicle” type in both truck & sedan. 
Give truck the field “fourWheel” which will be set to bool. 
Give sedan the field “luxury” which will be set to bool.

2) Using the vehicle, truck, and sedan structs: 
using a composite literal, create a value of type truck and assign values to the fields; 
using a composite literal, create a value of type sedan and assign values to the fields. 
Print out each of these values. Print out a single field from each of these values.

3) Give a method to both the “truck” and “sedan” types with the following signature
	transportationDevice() string
Have each func return a string saying what they do. 
Create a value of type truck and populate the fields. 
Create a value of type sedan and populate the fields. Call the method for each value. */

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
func (t truck) transportationDevice() string {
	return fmt.Sprintln("this is a", t.vehicle.color, "truck")
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintln("this is a", s.vehicle.color, "sedan")
}

func main (){
    myTruck := truck{ vehicle{4, "blue"}, true}
    mySedan := sedan { vehicle{2, "red"}, false}
    
    fmt.Println(myTruck, mySedan)

    fmt.Println(myTruck.color, mySedan.luxury)
    fmt.Println(myTruck.transportationDevice(),mySedan.transportationDevice())

}
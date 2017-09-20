// create a struct that holds person fields
// create a struct that holds secret agent fields and embeds person type
// attach a method to person: pSpeak
// attach a method to secret agent: saSpeak
// create a variable of type person
// create a variable of type secret agent
// print a field from person
// run pSpeak attached to the variable of type person
// print a field from secret agent
// run saSpeak attached to the variable of type secret agent
// run pSpeak attached to the variable of type secret agent

package main

import (
	"fmt"
)

type person struct {
	fName string
	lName string
}

type secretagent struct {
	person
	secretAgent bool
}

func (p person) pSpeak() string {
	return "Hi, I am " + p.fName
}
func (sa secretagent) saSpeak() string {
	return fmt.Sprintln("Hi, I know " + sa.fName)
}

func main() {
	person1 := person{"Jany", "Kim"}
	fmt.Println(person1.lName)

	sa1 := secretagent{person{"Tom", "Smith"}, true}
	fmt.Println(sa1.person.lName)
	fmt.Println(person1.pSpeak())
	fmt.Println(sa1.saSpeak())

}

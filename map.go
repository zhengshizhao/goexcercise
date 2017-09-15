/*Initialize a MAP using a composite literal where the key is a string and the value is an int; 
print out the map; range over the map printing out just the key; 
range over the map printing out both the key and the value */

package main

import ("fmt")

func main () {
	user := map[string]int {"age": 12, "zip": 10038}
	fmt.Println(user)
	for k := range user {
		println(k)
	}
	for k,v := range user {
		fmt.Println(k, v)
	}
	for i, c := range "go" {
        fmt.Println(i, c)
    }

}
/*Using the short declaration operator, create a variable with the identifier “s” 
and assign “i'm sorry dave i can't do that” to “s”. 
Print “s”. 
Print “s” converted to a slice of byte. 
Print “s” converted to a slice of byte and then converted back to a string. 
Using slicing, print just “i’m sorry dave”
Using slicing, print just “dave i can’t”
Using slicing, print just “can’t do that”
print every letter of “s” with one rune (character) on each line*/ 

package main 

import "fmt"

func main(){

	s :="i'm sorry dave i can't do that"
	fmt.Println(s)
	// b := make([]byte, len(s))
	// for i :=range s{
	// 	b[i] = byte(s[i])
	// }
	// fmt.Println(b)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))
	fmt.Println(s[:14])
	fmt.Println(s[10:22])
	fmt.Println(s[17:])
	for i := range s {
		fmt.Println(s[i])
	}


}
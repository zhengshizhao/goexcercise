/* Initialize a SLICE of int using a composite literal; 
print out the slice; range over the slice printing out just the index; 
range over the slice printing out both the index and the value */
package main

import ( "fmt"
)

func main () {
	thisWeek := []int{11, 12, 13, 14, 15}
	fmt.Println(`this week:`, thisWeek);
	for i := range thisWeek {
		fmt.Println(`index:`, i)
	} 

	for i, date := range thisWeek {
		fmt.Println(`index and value:`,i, `,`,date)
	}
	/*using make*/
    nextWeek := make([]int,5)
    for i := range thisWeek {
    	nextWeek[i] = thisWeek[i]+7
    }

    for i, date := range nextWeek {
		fmt.Println(`index and value:`,i, `,`,date)
	}
}

/*create a slice to store names,
show length,capacity,print out the values with index position without using range over.
*/
package main

import (
	"fmt"
)

func main() {
	x := make([]string, 5, 5)
	x = []string{"ani", "stephen", "emily", "george", "mark"}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

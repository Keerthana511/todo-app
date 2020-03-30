/*create a slice of slice of string
store data in multidimensional slice
*/
package main

import (
	"fmt"
)

func main() {
	xs1 := []string{"James", "Bond", "Shaken,not stirred"}
	xs2 := []string{"Miss", "Moneypenny", "Hello,James"}
	fmt.Println(xs1)
	fmt.Println(xs2)
	xxs := [][]string{xs1, xs2}
	fmt.Println(xxs)
	for i, xs := range xxs {
		fmt.Println("record:", i)
		for j, val := range xs {
			fmt.Printf("\tindex position :%v\tvalue:\t%v\n", j, val)
		}
	}
}

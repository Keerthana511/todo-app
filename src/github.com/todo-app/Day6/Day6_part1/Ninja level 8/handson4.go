/*sort the []int and []string.
 */
package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 6, 3, 2, 1, 8, 9, 13, 11}
	xs := []string{"James", "Q", "M", "Moneypenny", "Dr_no"}
	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("--------")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}

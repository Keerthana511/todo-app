/*fn foo that takes in variadic parameter of type int
unfurl the []int ,return the sum
fn bar that takes parameter of type []int
returns the sum of all values of type int passed in
*/
package main

import (
	"fmt"
)

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(ii...)
	fmt.Println(n)
	ii2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n1 := bar(ii2)
	fmt.Println(n1)

}
func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

/*create your own type that have an underlying type of struct.
 */
package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	favFlav []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlav: []string{
			"chocolate",
			"maritin",
			"rum and coke",
		},
	}
	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlav: []string{
			"strawberry",
			"vanila",
			"capuccino ",
		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFlav {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlav {
		fmt.Println(i, v)
	}
}

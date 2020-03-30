/*using prev code ,store the values of type in map
access each value,print the values,range over the slice
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
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.favFlav {
			fmt.Println(i, val)

		}
		fmt.Println("--------")

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

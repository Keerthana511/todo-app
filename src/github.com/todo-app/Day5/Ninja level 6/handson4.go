/*user defined struct,
attach a method,create a value of type,
call the method from the value
*/
package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}
func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}
	p1.speak()
}

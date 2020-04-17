/*method sets
 */
package main

import (
	"fmt"
)

type person struct {
	first string
}
type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("hello")
}
func saysomething(h human) {
	h.speak()
}
func main() {
	p1 := person{
		first: "James",
	}
	saysomething(&p1)
	p1.speak()
}

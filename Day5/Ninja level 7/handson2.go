/*changing the value stored at the addr,print the value.
 */
package main

import (
	"fmt"
)

type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond",
	}
	fmt.Println(p1)
	changeme(&p1)
	fmt.Println(p1)
}
func changeme(p *person) {
	p.name = "Miss Moneypenny"
	(*p).name = "Miss Money"
}

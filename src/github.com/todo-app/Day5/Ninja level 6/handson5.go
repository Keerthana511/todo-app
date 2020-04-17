/*create type,attach methods
create type which defines interface,which has the methods
create a func which takes type (interface) ,then prints.
create value and print .
*/
package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	area() float64
}

func info(s shape) {
	x := s.area()
	fmt.Println(x)
}
func main() {
	circ := circle{
		radius: 12.4,
	}
	squa := square{
		length: 13.5,
	}
	info(circ)
	info(squa)
}

/*create new types,underlying type struct
create a value of type and assign values to the fields
print out each of the values
print  out a single field from each of these values
*/
package main

import (
	"fmt"
)

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourwheel: true,
	}
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: false,
	}
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)

}

/*checking and handling error
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Hey", "Hi", "Heyloo"},
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln("Json did not Marshal here is the error:", err)
	}
	fmt.Println(string(bs))
}

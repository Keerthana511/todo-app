/*create a custom error msg using "fmt.Errorf"
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
	bs, err := toJSON(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("there was an error in to JSON:%v", err)
	}
	return bs, nil
}

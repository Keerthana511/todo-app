/*create a map key of type string  which is a persons last_first name
which records their fav 3 things
print the values along with index positions
*/
package main

import (
	"fmt"
)

func main() {
	m := map[string][]string{
		`bond_james`:  []string{`shaken,not stirred`, `martinis`, `women`},
		`money_penny`: []string{`James Bond`, `literature`, `cs`},
		`no_dr`:       []string{`being evil`, `icecream`, `sunsets`},
	}
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

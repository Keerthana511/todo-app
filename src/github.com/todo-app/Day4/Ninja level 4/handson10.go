/*using prev code,delete record from your map.
print map using range loop
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
    m[`fleming_ian`]=[]string{`steaks`, `eigars`, `espionage`}
    delete(m,`no_dr`)
    
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

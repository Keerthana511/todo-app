/*print every rune code point of the uppercase alphabets 3 times like,
U+0041 'A'
U+0041 'A'
U+0041 'A'
through the rest of the alphabet characters
*/
package main

import (
	"fmt"
)

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

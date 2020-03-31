/*create a value ,assign it to a variable.
print the addr of that value.
*/
package main
import(
    "fmt"
)
func main(){
    x:=42
    fmt.Println(x)
    fmt.Println(&x)
}
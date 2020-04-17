/*adding, benchmarks, examples, tests,coverage
 */
package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
func TestYears(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("got", n, "want", 70)
	}
}
func TestYearsTwo(t *testing.T) {
	n := Years(10)
	if n != 70 {
		t.Error("got", n, "want", 70)
	}
}
func ExampleYears() {
	fmt.Println(Years(10))
	//output:
	//70
}
func ExampleYearsTwo() {
	fmt.Println(Years(10))
	//output:
	//70
}

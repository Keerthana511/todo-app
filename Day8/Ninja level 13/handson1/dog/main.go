//package dog allows us to understand dogs
package dog

//years converts human years to dog years
func Years(n int) int {
	return n * 7
}

//yearsTwo converts human years to dog years
func YearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}

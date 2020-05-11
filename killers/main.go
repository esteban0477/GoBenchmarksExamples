package killers

/* func main() {
	fmt.Println("1+2+3 =", mySum(1, 2, 3))
	fmt.Println("1+2+3+4 =", mySum(1, 2, 3, 4))
	fmt.Println("1+2+3+4+5 =", mySum(1, 2, 3, 4, 5))
	fmt.Println("1+2+3+4+5+6 =", mySum(1, 2, 3, 4, 5, 6))
} */

// Sum func returns the sum of the integers passed as an argument
func Sum(intSlice ...int) int {
	result := 0
	for _, value := range intSlice {
		result += value
	}
	return result
}

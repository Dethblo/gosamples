package functions

import "fmt"

func DoVariadic() {
	numbers := []int{1, 2, 3, 4}
	fmt.Printf("sumup as an unpacked slice: %v\n", sumup(numbers...))
	fmt.Printf("sumup as separate values: %v\n", sumup(100, 200, 300))
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

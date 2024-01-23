package functions

import "fmt"

// DoMain the main caller of the examples.
func DoMain() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, getTransformerFn(false))
	tripled := transformNumbers(&numbers, getTransformerFn(true))
	squared := transformNumbers(&numbers, func(number int) int {
		return number * number
	})
	hundred := transformNumbers(&numbers, createTransformer(100))

	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(squared)
	fmt.Println(hundred)
}

// a type used as a shortcut for the function signature.
type transformFn func(int) int

func transformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func getTransformerFn(doTriple bool) transformFn {
	if doTriple {
		return triple
	} else {
		return double
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

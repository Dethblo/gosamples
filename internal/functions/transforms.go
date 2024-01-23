package functions

import "fmt"

type transformFn func(int) int

func TransformNumbers(numbers *[]int, transform transformFn) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func GetTransformerFn(doTriple bool) transformFn {
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

func CreateTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func DoMain() {
	numbers := []int{1, 2, 3, 4}
	doubled := TransformNumbers(&numbers, GetTransformerFn(false))
	tripled := TransformNumbers(&numbers, GetTransformerFn(true))
	squared := TransformNumbers(&numbers, func(number int) int {
		return number * number
	})
	hundred := TransformNumbers(&numbers, CreateTransformer(100))

	fmt.Println(doubled)
	fmt.Println(tripled)
	fmt.Println(squared)
	fmt.Println(hundred)
}

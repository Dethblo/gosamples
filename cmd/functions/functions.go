package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := transformNumbers(&numbers, getTransformerFn(false))
	tripled := transformNumbers(&numbers, getTransformerFn(true))

	fmt.Println(doubled)
	fmt.Println(tripled)
}

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

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)

	fmt.Println(doubled)
}

func doubleNumbers(numbers *[]int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, val*2)
	}
	return dNumbers
}

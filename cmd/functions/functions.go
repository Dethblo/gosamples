package main

import (
	"example.com/gosamples/internal/functions"
	"fmt"
)

func main() {
	functions.DoMain()
	fmt.Println("--------------------")
	functions.DoVariadic()
}

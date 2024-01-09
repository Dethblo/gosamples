package main

import "fmt"

type Product struct {
	Name        string
	Description string
	Price       float64
}

func main() {
	// a slice (dynamic array) of products
	products := []Product{
		{
			Name:        "Product 1",
			Description: "Product 1 description",
			Price:       10.99,
		},
		{
			Name:        "Product 2",
			Description: "Product 2 description",
			Price:       11.99,
		},
		{
			Name:        "Product 3",
			Description: "Product 3 description",
			Price:       12.99,
		},
	}

	fmt.Println(products)

	// a single new product
	newProduct := Product{
		Name:        "Product 4",
		Description: "Product 4 description",
		Price:       13.99,
	}

	// append the single product
	newList := append(products, newProduct)
	fmt.Println(newList)

	// a slice of new products
	specialProducts := []Product{
		{
			Name:        "Special Product 1",
			Description: "Special Product 1 description",
			Price:       20.99,
		},
		{
			Name:        "Special Product 2",
			Description: "Special Product 2 description",
			Price:       21.99,
		},
		{
			Name:        "Special Product 3",
			Description: "Special Product 3 description",
			Price:       22.99,
		},
	}

	// append multiple products to the list by unpacking the new list with '...' operator
	newList = append(newList, specialProducts...)
	fmt.Println(newList)

}

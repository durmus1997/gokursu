package main

import (
	"fmt"
	"golesson/project"
)

func main() {
	//variables.Demo1()
	//conditionals.Demo1()
	//conditionals.Demo2()
	//conditionals.WorkShop1()
	//arrays.Demo3()

	product, err := project.AddProduct()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(product)

	products, _ := project.GetAllProducts()

	for _, product := range products {
		fmt.Println(product.ProductName)
	}
}

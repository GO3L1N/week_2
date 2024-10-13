package main

import (
	"fmt"
	"main/product"
)

func main() {
	p := product.NewProduct("Laptop", 999.99)

	fmt.Println("Product Name:", p.GetName())

	// fmt.Println("Product Price:", p.price)  // это вызовет ошибку компиляции

	// fmt.Println("Product Price:", p.getPrice())  // это вызовет ошибку компиляции
}

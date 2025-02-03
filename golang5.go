package main

import (
	"fmt"
)

type taxPrice interface {
	calculateTaxPrice() float32
}

type costPrice interface {
	calculateCostPrice() float32
}

type Product struct {
	productName  string
	productPrice float32
	worthPercent float32
}

func (p Product) calculateTaxPrice() int {
	addTax := 12
	lastPrice := p.productPrice / (1 + float32(addTax)/100)
	lastPrice = p.productPrice - lastPrice
	return int(lastPrice)
}
func (p Product) calculateCostPrice() int {
	lastPrice := p.productPrice / (1 + float32(p.worthPercent)/100)
	lastPrice = p.productPrice - lastPrice
	return int(lastPrice)
}

func main() {

	product := Product{
		productName:  "Computer",
		productPrice: 2000,
		worthPercent: 20,
	}
	fmt.Println("Product Name", product.productName, "tax price", product.calculateTaxPrice())
	fmt.Println("Cost Price", product.calculateCostPrice())

}

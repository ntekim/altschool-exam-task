package main

import (
	"fmt"
)

type Store interface {
	addProduct() Product
	sellProduct() Product
	productList() []Product
	onSaleProducts() []Product
	soldProductList() []Product
	onSaleProductPriceTotal() float64
	soldProductPriceTotal() float64
	displayStatus() string
}
type Car struct {
	name  string
	model string
}
type Product struct {
	item       Car
	qtyInStock uint64
	price      float64
	outOfStock bool
}

var products []Product
var availableProducts = products
var soldProducts []Product

func (prod Product) addProduct() Product {

	products = append(products, prod)
	availableProducts = append(availableProducts, prod)
	fmt.Println("Product added successfully!")

	return prod
}

func (prod Product) sellProduct() Product {

	soldProducts = append(soldProducts, prod)
	availableProducts = products
	isEqual := false
	index := 0
	for i, item := range availableProducts {
		if item.item.name == prod.item.name {
			isEqual = true
			index = i
			break
		}
	}
	if isEqual {
		availableProducts = append(availableProducts[:index-1], availableProducts[index+1:]...)
	}
	fmt.Println("Sale successful!")

	return prod
}
func (prod Product) onSaleProductPriceTotal() float64 {
	var totalPrice float64
	for _, item := range soldProducts {
		if item.qtyInStock > 1 {
			totalPrice += item.price * float64(item.qtyInStock)
		} else {
			totalPrice += item.price
		}
	}

	return totalPrice
}

func (prod Product) displayStatus() string {
	if prod.outOfStock == false {
		return "Out of Stock!"
	} else {
		return "Still on Sale!"
	}
}

func (prod Product) soldProductPriceTotal() float64 {
	var totalPrice float64
	for _, item := range availableProducts {
		if item.qtyInStock > 1 {
			totalPrice += item.price * float64(item.qtyInStock)
		} else {
			totalPrice += item.price
		}
	}

	return totalPrice
}

func (prod Product) productList() []Product {

	return products
}

func (prod Product) soldProductList() []Product {

	return soldProducts
}

func (prod Product) onSaleProducts() []Product {

	return availableProducts
}

func main() {

	//
	var newStore Store
	//
	//var camry = Car{name: "Lexus", model: "360"}
	//
	//var newProd = Product{item: camry, price: 200000.00, qtyInStock: 10, outOfStock: false}

	newStore.addProduct()

	newStore.sellProduct()

	newStore.onSaleProducts()

	newStore.productList()

	newStore.soldProductList()
}

package controllers

import (
	"fmt"
	"github.com/ntekim/altschool-exam/models"
	"math/rand"
	"strconv"
)

type Product models.Product

type Store interface {
	AddProduct(prod models.Product) models.Product
	SellProduct(prod models.Product) models.Product
	ProductList() []models.Product
	OnSaleProducts() []models.Product
	SoldProductList() []models.Product
	OnSaleProductPriceTotal() float64
	SoldProductPriceTotal() float64
	DisplayStatus() string
}

var products []models.Product
var availableProducts = products
var soldProducts []models.Product

func (product Product) AddProduct(prod models.Product) models.Product {
	productExists := false
	index := 0
	for i, item := range availableProducts {
		if item.Item.Name == prod.Item.Name && item.Item.Model == prod.Item.Model {
			productExists = true
			index = i

			break
		}
	}
	if productExists {
		availableProducts[index].QtyInStock += 1
		availableProducts[index].TotalPrice += prod.Item.Price
	} else {
		prod.Id = strconv.Itoa(rand.Intn(1000000000))
		prod.TotalPrice = prod.Item.Price * float64(prod.QtyInStock)
		prod.OutOfStock = false
		products = append(products, prod)
		availableProducts = append(availableProducts, prod)
	}

	fmt.Println("Product added successfully!")

	return prod
}

func (product Product) SellProduct(prod models.Product) models.Product {

	//availableProducts := availProductList
	isEqual := false
	greaterThanOne := false
	index := 0
	for i, item := range availableProducts {
		if item.Item.Id == prod.Item.Id {
			isEqual = true
			index = i
			if item.QtyInStock > 1 {
				greaterThanOne = true
			}
			break
		}
	}
	if isEqual {
		if greaterThanOne {
			availableProducts[index].QtyInStock -= 1
			availableProducts[index].TotalPrice -= prod.Item.Price
			availableProducts[index].OutOfStock = false
			for _, item := range products {
				if item.Id == prod.Id {
					item.QtyInStock -= 1
					item.TotalPrice -= prod.Item.Price
					item.OutOfStock = false

					break
				}
			}
		} else {
			availableProducts = append(availableProducts[:index-1], availableProducts[index+1:]...)
			for _, item := range products {
				if item.Id == prod.Id {
					item.OutOfStock = true
					break
				}
			}
			productExists := false
			index := 0
			for i, item := range soldProducts {
				if item.Item.Name == prod.Item.Name && item.Item.Model == prod.Item.Model {
					productExists = true
					index = i

					break
				}
			}
			if productExists {
				soldProducts[index].QtyInStock += 1
				soldProducts[index].TotalPrice += prod.Item.Price
			} else {
				prod.TotalPrice = prod.Item.Price * float64(prod.QtyInStock)
				soldProducts = append(soldProducts, prod)
			}

		}

	}
	fmt.Println("Sale successful!")

	return prod
}
func (product Product) OnSaleProductPriceTotal() float64 {
	var totalPrice float64
	for _, item := range availableProducts {
		totalPrice += item.TotalPrice
	}

	return totalPrice
}

func (product Product) DisplayStatus() string {
	var status string
	for _, item := range products {
		if item.OutOfStock == false {
			status = "Still on sale!"
		} else {
			status = "Out of stock!"
		}
	}
	return status
}

func (product Product) SoldProductPriceTotal() float64 {
	var totalPrice float64
	for _, item := range soldProducts {
		totalPrice += item.TotalPrice
	}

	return totalPrice
}

func (product Product) ProductList() []models.Product {

	return products
}

func (product Product) SoldProductList() []models.Product {

	return soldProducts
}

func (product Product) OnSaleProducts() []models.Product {

	return availableProducts
}

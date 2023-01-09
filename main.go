package main

import (
	"github.com/ntekim/altschool-exam/controllers"
	"github.com/ntekim/altschool-exam/models"
)

func main() {

	var newStore controllers.Store

	var camry = models.Car{Name: "Lexus", Model: "360", Price: 2300.00, IsCarType: true, Color: "Blue"}

	var newProd = models.Product{Item: camry, TotalPrice: 200000.00, QtyInStock: 10, OutOfStock: false}

	newStore.AddProduct(newProd)
	//
	newStore.SellProduct(newProd)

	newStore.OnSaleProducts()

	newStore.ProductList()

	newStore.SoldProductList()

	newStore.DisplayStatus()

	newStore.SoldProductPriceTotal()

	newStore.OnSaleProductPriceTotal()
}

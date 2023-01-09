package models

type Cart struct {
	Id             string
	NumberOfItems  int
	TotalPrice     float64
	InStock        bool
	ProductsInCart []Product
}

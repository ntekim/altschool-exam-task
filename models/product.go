package models

type Product struct {
	Id         string
	Item       Car
	QtyInStock uint64
	TotalPrice float64
	OutOfStock bool
}

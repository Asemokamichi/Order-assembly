package main

type Product struct {
	productID   int
	productName string
}

type OrderProduct struct {
	orderID   int
	productID int
	quantity  int
}

type Box struct {
	boxID   int
	boxName string
}

type ProductBox struct {
	productID int
	boxID     int
	isMain    int
}

type Order struct {
	product       Product
	orderProduct  OrderProduct
	additionalBox string
}

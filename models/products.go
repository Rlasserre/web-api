package models

import (
	"web-api/db"
)

type Product struct {
	Name, Description string
	Id, Quantity      int
	Price             float64
}

func ScanProducts() []Product {
	db := db.Dbconnection()
	selectAllProducts, err := db.Query("Select * from products")
	if err != nil {
		panic(err.Error())
	}
	p := Product{}
	products := []Product{}
	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64
		err = selectAllProducts.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity
		products = append(products, p)
	}
	defer db.Close()
	return products
}

func Dbconnection() {
	panic("unimplemented")
}

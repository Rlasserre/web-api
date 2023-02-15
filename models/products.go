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
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()

	return products
}

func CreateNewProduct(name, description string, price float64, quantity int) {
	db := db.Dbconnection()

	insertIntoDb, err := db.Prepare("insert into products (name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insertIntoDb.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProduct(id string) {
	db := db.Dbconnection()
	productDelete, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}

	productDelete.Exec(id)

	defer db.Close()

}

func EditProduct(id string) Product {
	db := db.Dbconnection()

	productFromDb, err := db.Query("select * products where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	editedProduct := Product{}

	for productFromDb.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = productFromDb.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		editedProduct.Name = name
		editedProduct.Description = description
		editedProduct.Price = price
		editedProduct.Quantity = quantity
	}

	defer db.Close()
	return editedProduct
}

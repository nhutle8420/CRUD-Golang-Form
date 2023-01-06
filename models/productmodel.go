package models

import (
	"Go-crud/config"
	"Go-crud/entities"
	"database/sql"
	"fmt"
)

type ProductModel struct {
	conn *sql.DB
}

func NewProductModel() *ProductModel {
	connt, err := config.DBConnetion()
	if err != nil {
		panic(err)
	}
	return &ProductModel{
		conn: connt,
	}
}
func (p *ProductModel) FinAll() ([]entities.Product, error) {
	rows, err := p.conn.Query("SELECT * FROM product")
	if err != nil {
		fmt.Printf("err Finall")
		panic(err)
	}
	defer rows.Close()
	var products []entities.Product

	for rows.Next() {
		var pro entities.Product
		rows.Scan(&pro.ProductID, &pro.ProductName, &pro.ProductNumber, &pro.ProductPrice)
		products = append(products, pro)
	}

	return products, nil

}
func (p *ProductModel) Create(product entities.Product) bool {
	result, err := p.conn.Exec("insert into product (ProductName, ProductNumber, ProductPrice ) values(?,?,?)",
		product.ProductName, product.ProductNumber, product.ProductPrice)

	if err != nil {
		fmt.Println("err models", err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
func (p *ProductModel) Update(product entities.Product) error {

	_, err := p.conn.Exec(
		"update product set ProductName = ?, ProductNumber = ?, ProductPrice = ? where id = ?",
		product.ProductName, product.ProductNumber, product.ProductPrice, product.ProductID)

	if err != nil {
		return err
	}

	return nil
}
func (p *ProductModel) Find(id int64, pro *entities.Product) error {
	return p.conn.QueryRow("SELECT * FROM product WHERE ProductID  = ?", id).Scan(
		&pro.ProductID,
		&pro.ProductName,
		&pro.ProductNumber,
		&pro.ProductPrice)
}

func (p *ProductModel) Delete(id int64) {
	p.conn.Exec(" DELETE FROM product WHERE  ProductID  = ?", id)
}

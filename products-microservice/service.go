package main 

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-kit/kit/log"
)

type ProductsService interface {
	getProducts()([]*Product)
}

type DeleteService interface {
	deleteProduct(string, int64, bool) (bool)
}

type productsService struct{}
type deleteService struct{}

type Product struct {
	productID string
	name string
	quantity int64
	price float64
	image string
}

// Establishing connection with Mysql 
const serverName = "127.0.0.1:3306"
const dbName = "shop"
const user = "root"
const password = "drowssap"

// Business login for out /login RPC
func (productsService) getProducts() ([]*Product) {
	logger := log.NewLogfmtLogger(os.Stderr)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db ,_ := sql.Open("mysql", connectionString)
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		logger.Log(err)
	}
	defer rows.Close()
	defer db.Close()

	products := make([]*Product, 0)
	for rows.Next() {
		product := new(Product)
		err := rows.Scan(&product.productID, 
						&product.name, 
						&product.quantity, 
						&product.price, 
						&product.image)
		if err != nil {
			logger.Log(err)
		}
		products = append(products, product)
	}
	return products
}

// Business logic for /delete RPC
func (deleteService) deleteProduct(id string, quantity int64, admin bool) (bool) {
	logger := log.NewLogfmtLogger(os.Stderr)
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db ,_ := sql.Open("mysql", connectionString)

	defer db.Close()
	success := true
	if admin {
		_, err := db.Exec("DELETE FROM products WHERE productID = ?", id)
	
		if err != nil {
			logger.Log(err)
			return !success
		}
	} else {
		return !success
	}
	return success
}

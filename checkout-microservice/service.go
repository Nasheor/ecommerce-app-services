package main 

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)
type CheckoutService interface {
	checkout(int, int, int)(bool)
}

type checkoutService struct{}

type Checkout struct {
	CID int `json:customer_id`
	ID int `json:"product_id"`
	Q int `json:"quantity"`
}


// Establishing connection with Mysql 
const serverName = "127.0.0.1:3306"
const dbName = "shop"
const user = "root"
const password = "drowssap"

// Business login for out /login RPC
func (checkoutService) checkout(customer_id int, product_id int, quantity int) (bool) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db ,_ := sql.Open("mysql", connectionString)
	stmt, err := db.Prepare("INSERT INTO Orders(customerID, productID, quantity) VALUES( ?, ?, ? )")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close() // Prepared statements take up server resources and should be closed after use.
	if _, err := stmt.Exec(customer_id, product_id, quantity); err != nil {
		log.Fatal(err)
		return false;
	}
	return true
}
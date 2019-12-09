package main

import (
	"fmt"
	"os"
	"errors"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-kit/kit/log"
)

type LoginService interface {
	validateCredentials(string, string, bool)(bool)
}
type loginService struct{}

type User struct {
	U string `json:"u"`
	P string `json:"p"`
}

// Business login for out /login RPC
func (loginService) validateCredentials(uname string, pword string, admin bool) (bool) {
	logger := log.NewLogfmtLogger(os.Stderr)
	logger.Log(uname);
	serverName := "127.0.0.1:3306"
	dbName := "shop"
	user := "root"
	password := "drowssap"
	result := false

	// Establishing connection with Mysql 
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db ,err:= sql.Open("mysql", connectionString)
	if err != nil {
		logger.Log("Error in accessing the database")
	}
	defer db.Close()

	var u User

	// Checks the type of user that is customer or admin and queries the appropriate table
	if admin {
		err := db.QueryRow("SELECT name, password FROM admin where name=?", uname).Scan(&u.U, &u.P)
		if err != nil {
			logger.Log("Error reading from the database")
		}
	} else {
		err := db.QueryRow("SELECT name, password FROM Customer where name=?", uname).Scan(&u.U, &u.P)
		if err != nil {
			logger.Log("Error reading from the database")
		}
	}
	if u.P == pword && u.U == uname  {
		result = true

		
	}
	return result

}

// ErrEmpty is returned when an input string is empty
var ErrEmpty = errors.New("Empty Credentials")
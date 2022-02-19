package models

import (
	"fmt"
	"os"

	modelsBook "book-manager-api/models/book"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type dsn struct {
	user     string
	pass     string
	host     string
	port     string
	protocol string
	dbname   string
	options  map[string]string
}

var DB *gorm.DB

func ConnectDB() error {
	dsn := &dsn{
		user:     "root",
		pass:     os.Getenv("ROOT_PASS"),
		host:     "db",
		port:     "3306",
		protocol: "tcp",
		dbname:   os.Getenv("DATABASE"),
		options:  map[string]string{"parseTime": "True"},
	}

	dsnString := dsn.user + ":" + dsn.pass + "@" + dsn.protocol + "(" +
		dsn.host + ":" + dsn.port + ")/" + dsn.dbname

	// append options to dsnString
	if numOptions := len(dsn.options); numOptions != 0 {
		dsnString += "?"

		i := 1 // iterator
		for k, v := range dsn.options {
			dsnString += k + "=" + v
			if i != numOptions {
				dsnString += "&"
			}
			i++
		}
	}

	var err error
	if os.Getenv("DATABASE_URL") != "" {
		DB, err = gorm.Open(mysql.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})
	} else {
		DB, err = gorm.Open(mysql.Open(dsnString), &gorm.Config{})
	}

	if err != nil {
		return err
	}
	fmt.Println("DB connected: ", &DB)

	DB.AutoMigrate(&modelsBook.Book{})

	return nil
}

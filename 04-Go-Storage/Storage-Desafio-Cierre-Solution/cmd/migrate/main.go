package main

import (
	"app/internal/application"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &application.ConfigApplicationMigrate{
		Db: &mysql.Config{
			User:                 "root",
			Passwd:               "",
			Net:                  "tcp",
			Addr:                 "localhost:3306",
			DBName:               "fantasy_products",
		},
		FilePathCustomer: "./docs/db/json/customer.json",
		FilePathProduct: "./docs/db/json/product.json",
		FilePathInvoice: "./docs/db/json/invoice.json",
		FilePathSale: "./docs/db/json/sale.json",
	}
	app := application.NewApplicationMigrate(cfg)
	// - tear down
	defer app.TearDown()
	// - set up
	if err := app.SetUp(); err != nil {
		fmt.Println(err)
		return
	}
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
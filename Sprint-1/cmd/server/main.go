package main

import (
	"fmt"

	"github.com/usuario/repositorio/internal/application"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// app
	// - config
	addrCfg := ":8080"
	mysqlCfg := mysql.Config{
		User:      "melisprint_user",
		Passwd:    "melisprint_pass",
		Net:       "tcp",
		Addr:      "localhost:3306",
		DBName:    "melisprint",
		ParseTime: true,
	}
	cfg := application.ConfigServerChi{Addr: addrCfg, MySQLDSN: mysqlCfg.FormatDSN()}
	// - server
	server := application.NewServerChi(cfg)
	// - run
	if err := server.Run(); err != nil {
		fmt.Println(err)
		return
	}
}

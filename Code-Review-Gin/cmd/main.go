package main

import (
	"app/internal/application"
	"fmt"
)

func main() {
	// env
	// ...

	// app
	// - config
	cfg := &application.ConfigApplicationDefault{
		ServerAddress: ":8080",
		LoaderFilePath: "docs/db/vehicles_100.json",
	}
	app := application.NewApplicationDefault(cfg)
	// - setup
	err := app.SetUp()
	if err != nil {
		fmt.Println(err)
		return
	}
	// - run
	err = app.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}
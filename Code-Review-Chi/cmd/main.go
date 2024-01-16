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
	cfg := &application.ConfigServerChi{
		ServerAddress: ":8080",
		LoaderFilePath: "docs/db/vehicles_100.json",
	}
	app := application.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}
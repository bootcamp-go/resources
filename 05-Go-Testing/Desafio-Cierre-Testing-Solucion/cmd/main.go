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
		Addr: "127.0.0.1:8080",
	}
	app := application.NewApplicationDefault(cfg)
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
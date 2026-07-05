package main

import (
	"log"

	"github.com/11himanshusharma/memoria/internal/app"
	"github.com/11himanshusharma/memoria/internal/cli"
)

func main() {
	a, err := app.Bootstrap("configs/memoria.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := cli.Execute(a); err != nil {
		log.Fatal(err)
	}
}

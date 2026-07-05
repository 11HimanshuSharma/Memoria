package main

import (
	"github.com/11himanshusharma/memoria/internal/app"
	"github.com/11himanshusharma/memoria/internal/cli"
	"log"
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

package main

import (
	"log"
	"my-cli-app/app"
	"os"
)

func main() {
	if err := app.Gerar().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

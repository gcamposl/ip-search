package main

import (
	"fmt"
	"ip-search/app"
	"log"
	"os"
)

func main() {
	fmt.Println("initializing application...")
	application := app.Generate()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

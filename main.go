package main

import (
	"fmt"
	"ip-search/app"
	"log"
	"os"
)

func main() {
	fmt.Println("lookup ips and servers...")
	application := app.Generate()

	if err := application.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

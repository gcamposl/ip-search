package app

import "github.com/urfave/cli"

// generate returns the application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "search ips"
	app.Usage = "search ips and server names on the internet"

	return app
}

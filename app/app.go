package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

// generate returns the application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "search ips"
	app.Usage = "search ips and server names on the internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "search ips and addres on the internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

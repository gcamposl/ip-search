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

func searchServer(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host) // name server
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

// generate returns the application ready to run
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "search ips"
	app.Usage = "search ips and server names on the internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "search ips and addres on the internet",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "servers",
			Usage:  "search servers on the internet",
			Flags:  flags,
			Action: searchServer,
		},
	}

	return app
}

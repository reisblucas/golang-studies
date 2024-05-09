package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI Application"
	app.Usage = "Search for IPs and Server on internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search for iPs and Server on internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "google.com.br",
				},
			},
			Action: searchIPs,
		},
	}

	return app
}

func searchIPs(c *cli.Context) {
	host := c.String("host")

	// net package
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

package app

import (
	"github.com/urfave/cli"
	"log"
	"net"
)

// Build Will return CLI to be executed
func Build() *cli.App {
	app := cli.NewApp()

	app.Name = "CLI Application"
	app.Usage = "Search for IPs address"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IPs address",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:    "servers",
			Aliases: []string{"s"},
			Usage:   "Search for servers name",
			Flags:   flags,
			Action:  searchServers,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		log.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		log.Println(server)
	}
}

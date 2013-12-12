package main

import (
	"boshlite/configuration"
	"github.com/codegangsta/cli"
	util "boshlite/util"
	termcolor "boshlite/terminalcolor"
	"os"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "bosh-lite"
	app.Usage = "A command line tool to facilitate using Bosh-Lite deployment"
	app.Version = "1.0.0.alpha"
	app.Commands = []cli.Command{
		{
			Name:      "add-route",
			ShortName: "ar",
			Usage:     "add a system route to access the bosh-lite deployed vms",
			Action: func(c *cli.Context) {
				config := configuration.Default()

				fmt.Printf("Adding route for %s through Bosh lite gateway: %s \n",
					termcolor.Colorize(config.IpRange, termcolor.Yellow, true),
					termcolor.Colorize(config.Gateway, termcolor.Cyan, true))

        util.Addroute(config)
			},
		},
	}
	app.Run(os.Args)
}

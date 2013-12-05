package main

import (
	"boshlite/configuration"
	"github.com/codegangsta/cli"
	termcolor "boshlite/terminalcolor"
	"os"
	"log"
	"os/exec"
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

        var routecmd string
        switch config.OStype {
        case "Darwin":
           routecmd = "route delete -net " + config.IpRange + " " + config.Gateway + " > /dev/null 2>&1;"
           routecmd += "route add -net " + config.IpRange + " " + config.Gateway
        case "Linux":
           routecmd = "route add -net " + config.IpRange + " gw " + config.Gateway
        default:
            log.Fatal("Not supported OS detected")
        }

        out, err := exec.Command("sudo", "bash", "-c", routecmd).Output()
        if err != nil {
             log.Fatal(err)
        }
        fmt.Printf("%s", out)
			},
		},
	}
	app.Run(os.Args)
}

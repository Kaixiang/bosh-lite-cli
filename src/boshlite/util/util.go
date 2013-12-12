package util

import (
	"boshlite/configuration"
	"fmt"
	"log"
	"os/exec"
)

func Execute(bash string) ([]byte, error) {
	out, err := exec.Command("sudo", "bash", "-c", bash).Output()
	return out, err
}

func RouteCmd(config configuration.Configuration) (routecmd string) {
	switch config.OStype {
	case "Darwin":
		routecmd = "route delete -net " + config.IpRange + " " + config.Gateway + " > /dev/null 2>&1;"
		routecmd += "route add -net " + config.IpRange + " " + config.Gateway
	case "Linux":
		routecmd = "route add -net " + config.IpRange + " gw " + config.Gateway
	default:
		log.Fatal("Not supported OS detected")
	}
	return

}

func Addroute(config configuration.Configuration) {
	routecmd := RouteCmd(config)

	out, err := Execute(routecmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}
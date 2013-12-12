package util

import (
	"boshlite/configuration"
	"fmt"
	"log"
	"os/exec"
)

func Execute(bash string, sudo bool) (out []byte, err error) {
	if sudo {
		out, err = exec.Command("sudo", "bash", "-c", bash).Output()
	} else {
		out, err = exec.Command("bash", "-c", bash).Output()
	}
	return
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

	out, err := Execute(routecmd, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
}

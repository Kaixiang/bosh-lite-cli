package configuration

import (
	"os/exec"
)

type Configuration struct {
	Target  string
	IpRange string
	Gateway string
	OStype  string
	Version string
}

func DetectOS() string {
	out, err := exec.Command("bash", "-c", "uname").Output()
	if err != nil {
		return "UNKNOWN"
	}
	return string(out[:6])
}

func Default() (c Configuration) {
	c.Target = "http://api.10.244.0.34.xip.io"
	c.IpRange = "10.244.0.0/22"
	c.Gateway = "192.168.50.4"
	c.OStype = DetectOS()
	c.Version = "1"
	return
}

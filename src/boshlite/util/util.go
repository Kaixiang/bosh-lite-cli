package util

import (
  "fmt"
  "log"
  "os/exec"
  "boshlite/configuration"
)

func Execute(bash string) ([]byte, error) {
   out, err := exec.Command("sudo", "bash", "-c", bash).Output()
   return out, err
}

func Addroute(config configuration.Configuration) {
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

  out, err := Execute(routecmd)
  if err != nil {
       log.Fatal(err)
  }
  fmt.Printf("%s", out)
}

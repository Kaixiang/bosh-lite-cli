package util

import (
	"boshlite/configuration"
	termcolor "boshlite/terminalcolor"
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

type CheckVersion struct {
	name           string
	cmd_exist      string
	cmd_version    string
	expect_version string
}

func BuildSanitymap() []CheckVersion {
	check_map := []CheckVersion{
		{
			"Vagrant",
			"which vagrant",
			"vagrant -v|cut -d' ' -f2",
			"1.3.4",
		},
		{
			"Vagrant fusion plugin",
			"vagrant plugin list|grep vagrant-vmware-fusion",
			"vagrant plugin list|grep vagrant-vmware-fusion|cut -d' ' -f2|cut -d'(' -f2|cut -d')' -f1",
			"2.1.0",
		},
		{
			"BOSH CLI",
			"which bosh",
			"bosh -v|cut -d' ' -f2",
			"1.5.0.pre.1525",
		},
		{
			"Spiff",
			"which spiff",
			"spiff --version|cut -d' ' -f3",
			"0.0.0",
		},
		{
			"Go CF Cli",
			"which gcf",
			"gcf --version|cut -d' ' -f3",
			"6.0.0.rc1-SHA",
		},
		{
			"Virtual Box",
			"which VirtualBox",
			"VirtualBox --help|head -1|cut -d' ' -f 5",
			"4.2.18",
		},
	}
	return check_map
}

func SoftCheck() {
	sanity_map := BuildSanitymap()
	for _, check := range sanity_map {
		fmt.Printf("Checking %s...\n", termcolor.Colorize(check.name, termcolor.Cyan, false))
		_, err := Execute(check.cmd_exist, false)
		if err != nil {
			fmt.Printf("%s\n", termcolor.FailureColor("  [ERROR] No "+check.name+" found in your path"))
		} else {
			fmt.Printf("%s\n", termcolor.SuccessColor("  Found "+check.name+" installed"))
			out, err := Execute(check.cmd_version, false)
			if err != nil {
				fmt.Printf("%s\n", termcolor.FailureColor("  [ERROR] Not able to determine your "+check.name+" version"))
			}

			// create a string stip the newline
			var cur_version string
			if len(string(out)) > 1 {
				cur_version = string(out[:len(string(out))-1])
			} else {
				cur_version = "NIL"
			}
			if cur_version == "NIL" {
				fmt.Printf("%s\n", termcolor.WarnColor("  [Warnning] "+check.name+" version unknown, try install "+check.expect_version+" or newer"))
			} else if cur_version < check.expect_version {
				fmt.Printf("%s\n", termcolor.WarnColor("  [Warnning] Detect "+check.name+" version ("+cur_version+") lower than expected ("+check.expect_version+")"))
			} else {
				fmt.Printf("%s\n", termcolor.SuccessColor("  Detect "+check.name+" version ("+cur_version+") fulfill expected version ("+check.expect_version+")"))
			}
		}
	}
}

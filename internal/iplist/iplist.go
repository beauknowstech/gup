package iplist

import (
	"fmt"
	"net"
	"os"
	"regexp"
)

// https://github.com/mccoyst/myip/blob/master/myip.go
func LocalIP() {
	ipv4_regex := `^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			//https://gist.github.com/ahmetozer/ffa4cd0b319aff32ea9ed0068c8b81cf
			match, _ := regexp.MatchString(ipv4_regex, ipnet.IP.String())
			if match {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
}

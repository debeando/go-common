package net

import (
	"flag"
	"net"
	"os"
)

func Hostname() string {
	if flag.Lookup("test.v") != nil {
		return "localhost.test"
	}

	host, err := os.Hostname()
	if err != nil {
		return ""
	}

	return host
}

func IPAddress() string {
	if flag.Lookup("test.v") != nil {
		return "127.0.0.1"
	}

	addrs, _ := net.InterfaceAddrs()

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

package main

import (
	"fmt"
	"net"
)

func lookIP(address string) ([]string, error) {
	return net.LookupAddr(address)
}

func lookHostname(hostname string) ([]string, error) {
	return net.LookupHost(hostname)
}

func main() {
	ips, _ := lookHostname("localhost")
	for _, ip := range ips {
		fmt.Println(ip)
	}
	fmt.Println()
	hosts, _ := lookIP("127.0.0.1")
	for _, hostname := range hosts {
		fmt.Println(hostname)
	}
}

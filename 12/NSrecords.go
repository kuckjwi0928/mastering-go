package main

import (
	"fmt"
	"net"
)

func main() {
	domain := "google.com"
	NSs, _ := net.LookupNS(domain)
	for _, NS := range NSs {
		fmt.Println(NS.Host)
	}
}

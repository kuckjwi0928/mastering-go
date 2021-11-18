package main

import (
	"fmt"
	"net"
)

func main() {
	domain := "google.com"
	MXs, _ := net.LookupMX(domain)
	for _, MX := range MXs {
		fmt.Println(MX.Host)
		fmt.Println(MX.Pref)
	}
}

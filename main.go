package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	var CIDR = flag.String("cidr", "anonymous", "type cidr")

	flag.Parse()

	// generate a range of IP version 4 addresses from a Classless Inter-Domain Routing address
	ipAddress, ipNet, err := net.ParseCIDR(*CIDR)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// generate a range of IPv4 addresses from the CIDR address
	var ipAddresses []string

	for ipAddress := ipAddress.Mask(ipNet.Mask); ipNet.Contains(ipAddress); inc(ipAddress) {
		//fmt.Println(ipAddress)
		ipAddresses = append(ipAddresses, ipAddress.String())
	}

	// list out the ipAddresses within range
	for key, ipAddress := range ipAddresses {
		fmt.Printf("[%v] %s\n", key, ipAddress)
	}

}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

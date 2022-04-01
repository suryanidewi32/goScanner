package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {
	var IP = flag.String("ip", "", "type ip")
	var SUBNET = flag.Int("subnet", 1234, "subnet")
	flag.Parse()
	ip := *IP
	sub := *SUBNET
	wg.Add(sub)
	for i := 1; i <= sub; i++ {
		true_ip := ip + strconv.Itoa(i)
		go ping(true_ip)
	}
	wg.Wait()
}

func ping(ipAddress string) {
	var beaf = "false"
	Command := fmt.Sprintf("ping -c 1 %s  > /dev/null && echo true || echo false", ipAddress)
	output, err := exec.Command("/bin/sh", "-c", Command).Output()
	if err != nil {
		fmt.Println(err)
		return
	}
	real_ip := strings.TrimSpace(string(output))

	if real_ip == beaf {
		fmt.Printf("IP: %s   fail \n", ipAddress)
	} else {

		fmt.Printf("IP: %s   success ping pass \n", ipAddress)
	}
	wg.Done()
}

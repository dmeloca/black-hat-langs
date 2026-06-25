package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func scan(host string, port string, verbose bool) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second)
	if err != nil {
		if verbose {
			fmt.Printf("|- Port %s -> %v\n", port, err)
		}
	}
	if conn != nil {
		defer conn.Close()
		fmt.Printf("|- Open port on %s\n", port)
	}
}

func main() {
	fmt.Println("Ch1p's port scanner")

	end := flag.Int("p", 10000, "Number of ports to scan")
	verbose := flag.Bool("v", false, "Print more useful information")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("[!] Usage: scanner [options] <host>")
		os.Exit(1)
	}

	if *end < 1 {
		fmt.Println("[!] Port range must be greater than 0")
		os.Exit(1)
	}

	host := flag.Args()[0]

	fmt.Println("[!] Starting scan on", host)
	fmt.Println("[!] Ports to scan:", *end)

	for port := 1; port <= *end; port++ {
		scan(host, strconv.Itoa(port), *verbose)
	}
}

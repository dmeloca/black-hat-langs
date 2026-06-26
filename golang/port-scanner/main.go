package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func bannerGrab(conn net.Conn) string {
	conn.SetReadDeadline(time.Now().Add(time.Second))

	buffer := make([]byte, 1024)

	n, err := conn.Read(buffer)
	if err != nil {
		return "[-] No banner"
	}

	banner := string(buffer[:n])
	return banner
}

func scan(host string, port string, verbose bool) {
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), time.Second)

	if err != nil {
		if verbose {
			fmt.Printf("|- Port %s -> %v\n", port, err)
		}
		return
	}
	fmt.Printf("|- Open port on %s:%s\n", host, port)
	if port == "443" || port == "80" {
		return
	}
	defer conn.Close()
	fmt.Println("\t|- Banner: ", bannerGrab(conn))

}

func worker(host string, verbose bool, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for port := range jobs {
		scan(host, strconv.Itoa(port), verbose)
	}
}

func runWorkerPool(host string, end int, verbose bool) {
	const workerCount = 2000

	jobs := make(chan int)
	var wg sync.WaitGroup
	for range workerCount {
		wg.Add(1)
		go worker(host, verbose, jobs, &wg)
	}
	for port := 1; port <= end; port++ {
		jobs <- port
	}
	close(jobs)
	wg.Wait()
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

	if strings.Contains(host, "/") {

		fmt.Println("[!] Running scan in CIDR mode")

		ip, ipNet, err := net.ParseCIDR(host)
		if err != nil {
			if *verbose {
				fmt.Println(err)
			}
			os.Exit(1)
		}

		for ip := ip.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
			fmt.Printf("[!] Scanning %s\n", ip.String())
			runWorkerPool(ip.String(), *end, *verbose)
		}

	} else {

		fmt.Println("[!] Running scan in TARGET mode")
		runWorkerPool(host, *end, *verbose)

	}
}

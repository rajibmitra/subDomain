package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

const baseURL = "https://api.subdomain.center/?domain="

type Result struct {
	Subdomains []string
	Error      error
	Ports      map[string][]int
}

var wellKnownPorts = []int{
	20,    // FTP data
	21,    // FTP control
	22,    // SSH
	23,    // Telnet
	25,    // SMTP
	53,    // DNS
	80,    // HTTP
	110,   // POP3
	143,   // IMAP
	443,   // HTTPS
	465,   // SMTPS
	587,   // SMTP relay
	993,   // IMAPS
	995,   // POP3S
	3306,  // MySQL
	3389,  // RDP (Remote Desktop)
	5432,  // PostgreSQL
	27017, // MongoDB
	6379,  // Redis
	11211, // Memcached
	8080,  // HTTP alternate (often used for proxies or web servers)
	8443,  // HTTPS alternate (common for web applications using SSL)
}

func scanPorts(domain string, ch chan map[string][]int) {

	var openPorts []int
	const timeout = 500 * time.Millisecond

	for _, port := range wellKnownPorts {
		address := fmt.Sprintf("%s:%d", domain, port)
		conn, err := net.DialTimeout("tcp", address, timeout)
		if err == nil {
			openPorts = append(openPorts, port)
			conn.Close()
		}
	}
	ch <- map[string][]int{domain: openPorts}
}

func fetchSubdomains(domain string, ch chan Result) {
	client := &http.Client{}

	resp, err := client.Get(baseURL + domain)
	if err != nil {
		ch <- Result{Error: err}
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		ch <- Result{Error: fmt.Errorf("API returned non-OK status: %v", resp.Status)}
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- Result{Error: err}
		return
	}

	var subdomains []string
	err = json.Unmarshal(body, &subdomains)
	if err != nil {
		ch <- Result{Error: err}
		return
	}

	if len(subdomains) == 0 {
		ch <- Result{Error: fmt.Errorf("no subdomains found for %v", domain)}
		return
	}
	portsMapping := make(map[string][]int)
	portCh := make(chan map[string][]int)

	for _, subdomain := range subdomains {
		go scanPorts(subdomain, portCh)
	}

	for range subdomains {
		portData := <-portCh
		for k, v := range portData {
			if len(v) > 0 {
				portsMapping[k] = v
			}
		}
	}

	ch <- Result{Subdomains: subdomains, Ports: portsMapping}
}

func main() {

	var domain string
	outputFormat := flag.String("o", "", "Output format (txt for text file)")
	flag.StringVar(&domain, "u", "", "Domain to fetch subdomains for")
	flag.Parse()

	if domain == "" {
		log.Fatalf("You must provide a domain with the -u flag.")
	}

	ch := make(chan Result)

	go fetchSubdomains(domain, ch)

	result := <-ch

	if result.Error != nil {
		log.Fatalf("Error fetching subdomains: %v", result.Error)
	}

	if *outputFormat == "txt" {
		var content strings.Builder
		for subdomain, ports := range result.Ports {
			if len(ports) > 0 {
				content.WriteString(fmt.Sprintf("%s: %v\n", subdomain, ports))
			} else {
				content.WriteString(fmt.Sprintf("%s\n", subdomain))
			}
		}
		err := os.WriteFile("subdomains.txt", []byte(content.String()), 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
		fmt.Println("Subdomains written to subdomains.txt")
	} else {
		for subdomain, ports := range result.Ports {
			if len(ports) > 0 {
				fmt.Printf("%s: %v\n", subdomain, ports)
			} else {
				fmt.Println(subdomain)
			}
		}
	}

}

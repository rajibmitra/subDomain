package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const baseURL = "https://api.subdomain.center/?domain="

type Result struct {
	Subdomains []string
	Error      error
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
		ch <- Result{Error: fmt.Errorf("No subdomains found for %v", domain)}
		return
	}

	ch <- Result{Subdomains: subdomains}
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
		content := strings.Join(result.Subdomains, "\n")
		err := os.WriteFile("subdomains.txt", []byte(content), 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
		fmt.Println("Subdomains written to subdomains.txt")
	} else {
		for _, subdomain := range result.Subdomains {
			fmt.Println(subdomain)
		}
	}
}

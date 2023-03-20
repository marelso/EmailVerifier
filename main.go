package main

import (
	"strings"
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarc \n")

	for scanner.Scan() {
		check(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read from input: %v \n", err)
	}
}

func check(domain string) {
	var hasMX, hasMX, hasDMARC bool
	var sprRecord, dmarcRecord string

	hasMX = validateMX(domain)

	txtRecords, err := net.LookupTXT(domain)


}

func validateMX(domain string) (bool) {
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error: %v \n", err)
	}

	if len(mxRecords) > 0 {
		return true
	}

	return false
}
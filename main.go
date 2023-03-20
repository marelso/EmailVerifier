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
	hasSPF, sprRecord = validateTXT(domain, false)
	hasDMARC, dmarcRecord = validateTXT(domain, true)

	fmt.Printf("%v, %v, %v, %v, %v, %v \n", domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarc)
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

func validateTXT(domain string, isDmarc bool) (bool, string) {
	txtRecords, err := net.LookupTXT(domain)

	var theRecord string
	var prefix string
	var has bool

	switch isDmarc {
		case true:
			prefix = "v=spf1":
			break

		case false:
			prefix = "v=DMARC1":
			break
	}

	if err != nil { 
		log.Printf("Error: %v", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, prefix) {
			has = true
			theRecord = record
		}
	}

	return has, theRecord
}
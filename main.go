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

	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord \n")

	if len(os.Args) >= 2 {
		for _, arg := range os.Args[1:] {
			check(arg)
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		check(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Could not read from input: %v \n", err)
	}
}

func check(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var sprRecord, dmarcRecord string

	hasMX = validateMX(domain)
	hasSPF, sprRecord = validateTXT(domain, false)
	hasDMARC, dmarcRecord = validateTXT(domain, true)

	fmt.Printf("%v, %v, %v, %v, %v, %v \n", domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord)
}

func validateMX(domain string) bool {
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
	var theRecord string
	var prefix string
	var has bool

	switch isDmarc {
		case true:
			prefix = "v=DMARC1"
		case false:
			prefix = "v=spf1"
	}

	txtRecords, err := net.LookupTXT(domain)

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
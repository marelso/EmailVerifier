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
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarc")

	for scanner.Scan() {
		check(scanner.Text())
	}
}

func check(domain string) {
	
}
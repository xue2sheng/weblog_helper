package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
)

// https://github.com/xue2sheng/weblog_helper

// DefaultFilePath default file path
const DefaultFilePath = "./public_access.log.txt"

// SepToken to split header words
const SepToken = " - - "

// matching IP or CIDR range
type feature func(string, string, *net.IPNet) bool

// Feature 1: just one ip
func featureIP(simpleIP string, ipMatch string, cidrMatch *net.IPNet) bool {
	// ignore cidrMatch
	return len(simpleIP) > 0 && simpleIP == ipMatch
}

// Feature 2: CIDR range
func featureCidr(simpleIP string, ipMatch string, cidrMatch *net.IPNet) bool {
	//ipAddress, IPnet, err := net.ParseCIDR("198.162.0.0/16")
	// ignore ipMatch

	// is an IP non empty
	if len(simpleIP) == 0 {
		return false
	}

	// is really an IP
	ip := net.ParseIP(simpleIP)
	if ip == nil {
		return false
	}

	// belong to taht CIDR???
	return cidrMatch.Contains(ip)
}

// 31.184.238.128 - - [02/Jun/2015:17:00:12 -0
func ipCandidate(firstItem *regexp.Regexp, line string, filter feature, ipMatch string, cidrMatch *net.IPNet) (candidate string, err error) {

	// get matches
	candidate = firstItem.Split(line, -1)[0]
	if len(candidate) == 0 || len(candidate) == len(line) {
		return "", errors.New("Unexpected header in line: " + line)
	}

	if filter(candidate, ipMatch, cidrMatch) {
		return line, nil
	}

	return "", errors.New("Filter failed for line: " + line)
}

// Only looking for IPv4 due to the fact that input file seems to prefer that version
func main() {

	// commandline arguments
	filename := flag.String("filename", DefaultFilePath, "File to parse. Default "+DefaultFilePath)
	matchItem := flag.String("ip", "", "MatchItem as IP or CIDR to look for. (Required)")
	flag.Parse()

	// needed something to look for
	if *matchItem == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// tow posibilities
	var featureItem feature
	var cidrMatch *net.IPNet
	var err error

	// Is a valid IP?
	trial := net.ParseIP(*matchItem)
	if trial.To4() == nil {

		// Maybe is a valid CIRR
		_, cidrMatch, err = net.ParseCIDR(*matchItem)
		if err != nil {
			log.Fatal("Invalid IPv4 address or CIDR match:" + *matchItem)
		}
		featureItem = featureCidr

	} else {
		featureItem = featureIP
	}

	// open file
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// scan for matches
	firstItem := regexp.MustCompile(SepToken)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		candidate, err := ipCandidate(firstItem, scanner.Text(), featureItem, *matchItem, cidrMatch)
		if err == nil {
			fmt.Println(candidate)
		}
	}

	// check errors
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"net"
	"regexp"
	"testing"
)

// IP0 default match
const IP0 = "31.184.238.128"

// CIDR0 range default match
const Cidr = "180.76.15.0/24"
const CirdOk = "180.76.15.135"
const CirdFail = "192.76.15.135"

func TestIsCandidateFeature1(t *testing.T) {

	firstItem := regexp.MustCompile(SepToken)
	str, err := ipCandidate(firstItem, IP0+" - - two three", featureIP, IP0, nil)
	if err != nil {
		t.Fatal(err)
	}
	if str != IP0+" - - two three" {
		t.Fatal("Not the header word: ", str)
	}

	str, err = ipCandidate(firstItem, "nothing to do", featureIP, IP0, nil)
	if err == nil {
		t.Fatal("That should have failed: ", str)
	}
}

func TestIsCandidateFeature2(t *testing.T) {

	firstItem := regexp.MustCompile(SepToken)
	_, cidrMatch, _ := net.ParseCIDR(Cidr)

	str, err := ipCandidate(firstItem, CirdOk+" - - two three", featureCidr, "", cidrMatch)
	if err != nil {
		t.Fatal(err)
	}
	if str != CirdOk+" - - two three" {
		t.Fatal("Not expected CIDR OK: ", str)
	}

	str, err = ipCandidate(firstItem, CirdFail+" - - to do", featureCidr, "", cidrMatch)
	if err == nil {
		t.Fatal("That should have failed: ", str)
	}
}

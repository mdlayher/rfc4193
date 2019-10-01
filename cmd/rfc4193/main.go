// Command rfc4193 generates a Unique Local IPv6 Unicast Address prefix, as
// described in RFC4193.
package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mdlayher/rfc4193"
)

func main() {
	ll := log.New(os.Stderr, "", 0)

	ifis, err := net.Interfaces()
	if err != nil {
		ll.Fatalf("failed to get network interfaces: %v", err)
	}

	// Try to choose a suitable interface MAC address as a seed, but also fall
	// back to random data (nil mac input) if a suitable address isn't found.
	var mac net.HardwareAddr
	for _, ifi := range ifis {
		// Must be Ethernet address, must be non-zero (skip loopback).
		if len(ifi.HardwareAddr) != 6 || bytes.Equal(ifi.HardwareAddr, make([]byte, 6)) {
			continue
		}

		mac = ifi.HardwareAddr
		break
	}

	p, err := rfc4193.Generate(mac)
	if err != nil {
		ll.Fatalf("failed to generate RFC4193 prefix: %v", err)
	}

	fmt.Println(p)
}

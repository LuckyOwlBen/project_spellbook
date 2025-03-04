package services

import (
	"fmt"

	"github.com/t94j0/nmap"
)

func MapNetwork(ipAddress string, lowIp uint16, highIp uint16) []nmap.Port {
	fmt.Println("Mapping network at ip", ipAddress)
	scanResult, error := nmap.
		Init().
		AddHosts(ipAddress).
		AddPortRange(lowIp, highIp).
		SetFlags("-sV", "-sC").
		Run()

	if error != nil {
		panic(error)
	}

	host, exists := scanResult.GetHost(ipAddress)

	if !exists {
		panic("Host not found")
	}

	return host.Ports

}

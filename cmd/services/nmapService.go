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
	fmt.Println("Scan result:", scanResult)
	host, exists := scanResult.GetHost(ipAddress)
	//fmt.Println("Host exists:", exists)
	//fmt.Println("Host:", host)

	if !exists {
		fmt.Println("Host not found")
		//panic("Host not found")
	}

	return host.Ports

}

package definitions

import (
	"encoding/xml"
	"fmt"
	"project_spellbook/utils"

	"github.com/lair-framework/go-nmap"
)

func RunNmap(ipAddress string) {
	args := []string{"nmap", "-oX", "-", "-sV", "-sC", ipAddress}
	nmapResult := utils.ExecuteCommand(true, args...)
	//println(string(nmapResult))

	var result nmap.NmapRun
	err := xml.Unmarshal(nmapResult, &result)
	if err != nil {
		fmt.Println("Error unmarshalling nmap result:", err)
		return
	}

	// Print the parsed results
	for _, host := range result.Hosts {
		fmt.Printf("Host: %s\n", host.Addresses[0].Addr)
		for _, port := range host.Ports {
			fmt.Printf("Port: %d, State: %s, Service: %s\n", port.PortId, port.State.State, port.Service.Name)
			if port.Service.Name == "telnet" {
				fmt.Printf("Telnet service found on port %d\n", port.PortId)
				ProbeTelnet(ipAddress)
			}
			if port.Service.Name == "ftp" {
				fmt.Printf("FTP service found on port %d\n", port.PortId)
				ProbeFTP(ipAddress)
			}
		}
	}
}

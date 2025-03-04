package main

import (
	"bufio"
	"fmt"
	"os"
	"project_spellbook/cmd/services"
	"strings"
)

func main() {
	services.ConnectToOpenVpn()
	fmt.Println("Provide the IP address you want to ping:")
	reader := bufio.NewReader(os.Stdin)
	ipAddress, _ := reader.ReadString('\n')
	ipAddress = strings.TrimSpace(ipAddress)
	fmt.Println("Pinging", ipAddress)
	services.Ping(ipAddress)
	ports := services.CompileNetworkMap(ipAddress)
	services.CheckForExploits(ports, ipAddress)
	services.DisconnectFromOpenVpn()
}

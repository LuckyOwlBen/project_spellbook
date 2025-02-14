package main

import (
	"bufio"
	"fmt"
	"os"
	"project_spellbook/definitions"
	"strings"
)

func main() {
	definitions.ConnectToOpenVpn()
	fmt.Println("Provide the IP address you want to ping:")
	reader := bufio.NewReader(os.Stdin)
	ipAddress, _ := reader.ReadString('\n')
	ipAddress = strings.TrimSpace(ipAddress)
	fmt.Println("Pinging", ipAddress)
	definitions.Ping(ipAddress)
	definitions.RunNmap(ipAddress)
	definitions.DisconnectFromOpenVpn()
}

package services

import (
	"fmt"
	"project_spellbook/cmd/console"
)

func Ping(ipAddress string) {
	args := []string{"ping", "-c", "4", ipAddress}
	pingResult := console.ExecuteCommand(false, args...)
	fmt.Println(string(pingResult))
}

package definitions

import (
	"fmt"
	"project_spellbook/utils"
	"strings"
)

func Ping(ipAddress string) {
	ipAddress = strings.TrimSpace(ipAddress)
	args := []string{"ping", "-c", "4", ipAddress}
	pingResult := utils.ExecuteCommand(false, args...)
	fmt.Println(string(pingResult))
}

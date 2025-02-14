package definitions

import (
	"fmt"
	"project_spellbook/utils"
)

func Ping(ipAddress string) {
	args := []string{"ping", "-c", "4", ipAddress}
	pingResult := utils.ExecuteCommand(false, args...)
	fmt.Println(string(pingResult))
}
package definitions

import (
	"fmt"
	"project_spellbook/utils"
)

func ConnectToOpenVpn() {
	args := []string{"openvpn", "/home/demeter/Downloads/starting_point_cyberblade9876.ovpn"}
	utils.ExecuteAsyncCommand(true, args...)
}

func DisconnectFromOpenVpn() {
	args := []string{"killall", "openvpn"}
	disconnectResult := utils.ExecuteCommand(true, args...)
	fmt.Println(string(disconnectResult))
}

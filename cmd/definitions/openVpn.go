package definitions

import (
	"fmt"
	"project_spellbook/cmd/console"
)

func ConnectToOpenVpn() {
	args := []string{"openvpn", "/home/demeter/Downloads/starting_point_cyberblade9876.ovpn"}
	console.ExecuteAsyncCommand(true, args...)
}

func DisconnectFromOpenVpn() {
	args := []string{"killall", "openvpn"}
	disconnectResult := console.ExecuteCommand(true, args...)
	fmt.Println(string(disconnectResult))
}

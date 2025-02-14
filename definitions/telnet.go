package definitions

import (
	"fmt"
	"strings"

	"github.com/reiver/go-telnet"
)

func ProbeTelnet(ipAddress string) {
	usernames := []string{"root", "admin", "user"}
	for _, username := range usernames {
		fmt.Printf("Trying to log in with username: %s\n", username)
		address := fmt.Sprintf("%s:23", ipAddress) // Format the IP address and port number
		conn, err := telnet.DialTo(address)        // Establish a Telnet connection to the address
		if err != nil {
			fmt.Println("Error connecting to Telnet:", err)
			continue
		}
		fmt.Println("Connected to Telnet")
		// Send the username
		_, err = conn.Write([]byte(username + "\n"))
		if err != nil {
			fmt.Println("Error sending username:", err)
			conn.Close()
			continue
		}

		// Read the response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response:", err)
			conn.Close()
			continue
		}

		response := string(buffer[:n])
		fmt.Println(response)

		if strings.Contains(response, "Welcome") {
			fmt.Println("Telnet successful with username", username)
			// Execute a command after successful login
			_, err = conn.Write([]byte("ls -la\n"))
			if err != nil {
				fmt.Println("Error sending command:", err)
				conn.Close()
				continue
			}

			// Read the command output
			n, err = conn.Read(buffer)
			if err != nil {
				fmt.Println("Error reading command output:", err)
				conn.Close()
				continue
			}

			fmt.Println(string(buffer[:n]))
			conn.Close()
			return
		}

		conn.Close()
	}
	fmt.Println("Telnet failed to connect to", ipAddress)
}

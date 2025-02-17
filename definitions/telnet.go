package definitions

import (
	"fmt"
	"strings"
	"time"

	"github.com/ziutek/telnet"
)

func ProbeTelnet(ipAddress string) {
	usernames := []string{"admin", "user", "root"}
	address := fmt.Sprintf("%s:23", ipAddress) // Format the IP address and port number
	fmt.Println("Connecting to Telnet", address)
	conn, err := telnet.Dial("tcp", address) // Establish a Telnet connection to the address
	if err != nil {
		fmt.Println("Connection Failed:", err)
		return
	}
	fmt.Println("Connected to Telnet", address)
	scrollScanBuffer(conn)
	for _, username := range usernames {
		fmt.Printf("Trying to log in with username: %s\n", username)
		// Send the username
		conn.SetUnixWriteMode(true)
		conn.Write([]byte(username + "\n"))
		conn.Write([]byte("\n"))     // Send the password
		time.Sleep(10 * time.Second) // Wait for the server to process the login
		// Read the response from the server

		buf := make([]byte, 50000)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read response:", err)
			break
		}

		response := string(buf[:n])
		fmt.Println(
			"Response:", response,
		)

		// Check if the response indicates a successful login
		if strings.Contains(response, "Welcome") || strings.Contains(response, "Success") {
			fmt.Println("Login succeeded with username:", username)
		} else {
			fmt.Println("Login failed with username:", username)
		}
	}
	conn.Close() // Close the connection
	fmt.Println("Connection closed")
}

func scrollScanBuffer(connection telnet.Conn) {
	buf := make([]byte, 50000)
	for {
		n, err := connection.Read(buf)
		if err != nil {
			fmt.Println("Failed to read response:", err)
			return
		}
		response := string(buf[:n])
		if strings.Contains(response, "login") {
			return
		}
	}

}

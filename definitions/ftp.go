package definitions

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jlaffaye/ftp"
)

func ProbeFTP(ipAddress string) {
	// Connect to the FTP server
	client, err := ftp.Dial(ipAddress+":21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}

	// Login to the FTP server
	err = client.Login("anonymous", "anonymous")
	if err != nil {
		log.Fatal(err)
	}

	// List the contents of the current directory
	entries, err := client.List("")
	if err != nil {
		log.Fatal(err)
	}

	// Print the contents of the current directory
	fmt.Println("Contents of the current directory:")
	for _, entry := range entries {
		fmt.Println(entry.Name)
		if entry.Name == "flag.txt" {
			fmt.Println("Flag found in FTP server")
			//get the flag
			buf, err := client.Retr("flag.txt")
			if err != nil {
				log.Fatal(err)
			}
			content, error := io.ReadAll(buf)
			if error != nil {
				log.Fatal(error)
			}
			fmt.Println(string(content))
		}
	}

	// Close the connection
	err = client.Quit()
	if err != nil {
		log.Fatal(err)
	}

}

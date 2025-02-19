package definitions

import (
	"fmt"
	"net"

	"github.com/hirochachacha/go-smb2"
)

func ProbeSMB(ipAddress string) {

	conn, err := net.Dial("tcp", ipAddress+":445")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     "fakey",
			Password: "",
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	defer s.Logoff()

	names, err := s.ListSharenames()
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		fmt.Println(name)
		share, err := s.Mount(name)
		if err != nil {
			fmt.Println("Access denied to share", err)
			continue
		}
		files, err := share.ReadDir("")
		if err != nil {
			fmt.Println("Error reading directory", err)
			continue
		}
		for _, directory := range files {
			fmt.Println(directory.Name())
			if !directory.IsDir() {
				fmt.Println("Skipping file", directory.Name())
				continue
			}
			subFiles, err := share.ReadDir(directory.Name())
			if err != nil {
				fmt.Println("Error scanning directory", err)
				continue
			}
			for _, file := range subFiles {
				fmt.Println(file.Name())
				if file.Name() == "flag.txt" {
					fmt.Println("Flag found in SMB server")
					fileContent, err := share.Open(directory.Name() + "/" + file.Name())
					if err != nil {
						fmt.Println("Error opening file", err)
						continue
					}
					defer fileContent.Close()
					content := make([]byte, file.Size())
					_, err = fileContent.Read(content)
					if err != nil {
						fmt.Println("Error reading file", err)
						continue
					}
					fmt.Println(string(content))
				}
			}
		}

	}
}

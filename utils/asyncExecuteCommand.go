package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteAsyncCommand(requiresSudo bool, args ...string) {
	if requiresSudo {
		args = append([]string{"sudo"}, args...)
	}

	fmt.Println(args)
	cmd := exec.Command(args[0], args[1:]...)
	//allows input of sudo password
	//cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

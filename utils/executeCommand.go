package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteCommand(requiresSudo bool, args ...string) string {
	if requiresSudo {
		args = append([]string{"sudo"}, args...)
	}

	fmt.Println("args: ", args)
	cmd := exec.Command(args[0], args[1:]...)
	fmt.Println("cmd: ", cmd)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err, "\n", string(output))
		return "error"
	}
	return string(output)
}

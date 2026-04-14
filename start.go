package main

import (
	"os/exec"
	"fmt"
)



func main() {
	programs := []string{"webpage", "api", "assets"}
	for _, v := range programs {
		cmd := exec.Command("go", "run", v + ".go")

		// Capture the output
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		}

		fmt.Printf("Output:\n%s\n", string(output))
	}
}
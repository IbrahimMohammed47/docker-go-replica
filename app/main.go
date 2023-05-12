package main

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
	// Uncomment this block to pass the first stage!
	// "os"
	// "os/exec"
)

// Usage: your_docker.sh run <image> <command> <arg1> <arg2> ...
func main() {

	mainCommand := os.Args[3]
	args := os.Args[4:len(os.Args)]

	cmd := exec.Command(mainCommand, args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// cmd.Run()
	err := cmd.Run()
	if err != nil && strings.HasPrefix(err.Error(), "exit status") {
		status := strings.SplitAfter(err.Error(), "exit status ")[1]
		// convert string to int
		statusInt, err2 := strconv.Atoi(status)
		if err2 == nil {
			os.Exit(statusInt)
		}
		os.Exit(120)
	}
}

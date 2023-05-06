package main

import (
	"os"
	"os/exec"
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

	cmd.Run()
	// output, err := cmd.Output()
	// if err != nil {
	// 	fmt.Print(err)
	// 	os.Exit(1)
	// }
	// out := make([]byte, 6)
	// cmd.Stdout.Write(out)
	// fmt.Print(string(output))
}

package main

import (
		"fmt"
		"io"
		"os/exec"
		"strings"
	)

func runCommandWithInput(cmdName string, args string, input string) string {
	cmd := exec.Command(cmdName, args)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return ""
	}

	go func() {
		defer stdin.Close()
		io.WriteString(stdin, input)
	}()

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		return ""
	}

	return string(stdout)
}

func getCurrentVersion() string {
	cmd := exec.Command("git", "tag", "--sort=committerdate")
	stdout, err := cmd.Output()

    if err != nil {
        return ""
    }

    // Print the output
    tag := runCommandWithInput("tail", "-n 1", string(stdout))
    return strings.Trim(tag, "\n")
}

func main() {
    fmt.Println("Hello, I'm ocp-project-api service!")

	fmt.Println("My version: ", getCurrentVersion())
}

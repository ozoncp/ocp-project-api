package main

import (
		"fmt"
		"io"
		"os/exec"
	)

func getCurrentVersion() string {
	cmd := exec.Command("git", "tag", "--sort=committerdate")
	stdout, err := cmd.Output()

    if err != nil {
        return ""
    }

    // Print the output
    tags := string(stdout)

	cmd = exec.Command("tr", "-n 1")
    stdin, err := cmd.StdinPipe()
    if err != nil {
        return ""
    }

    go func() {
        defer stdin.Close()
        io.WriteString(stdin, tags)
    }()

    stdout, err = cmd.CombinedOutput()
    if err != nil {
        return ""
    }

    return string(stdout)
}

func main() {
    fmt.Println("Hello, I'm ocp-project-api service!")

	fmt.Println("My version: ", getCurrentVersion())
}

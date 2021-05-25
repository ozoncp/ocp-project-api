package main

import (
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/utils"
)

func main() {
	fmt.Println("Hello, I'm ocp-project-api service!")
	utils.LoopOpenClose("test.txt", "Write some msg", 10)
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("command not found")
		}
		if strings.TrimRight(command, "\n") == "exit" {
			break
		} else if strings.HasPrefix(command, "echo ") {
			fmt.Print(command[5:])
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}

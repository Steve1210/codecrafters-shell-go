package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

var builtins = [...]string{"echo", "type", "exit"}

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
		} else if strings.HasPrefix(command, "type ") {
			arg := strings.TrimRight(command[5:], "\n")
			found := false
			for _, builtin := range builtins {
				if arg == builtin {
					fmt.Printf("%s is a shell builtin\n", arg)
					found = true
					break
				}
			}
			if !found {
				// if not builtin, Look for exec in PATH
				path, err := exec.LookPath(arg)
				if err != nil {
					fmt.Printf("%s: not found\n", arg)
				} else {
					fmt.Printf("%s is %s\n", arg, path)
				}
			}
		} else {
			fmt.Println(command[:len(command)-1] + ": command not found")
		}
	}
}

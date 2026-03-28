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
			// check if command is in PATH
			parts := strings.Split(strings.TrimRight(command, "\n"), " ")
			file := parts[0]
			_, err := exec.LookPath(file)
			if err != nil {
				fmt.Println(file + ": command not found")
			} else {
				//args := strings.Join(strings.Split(command, " ")[1:], "")
				cmd := exec.Command(file, parts[1:]...)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				if err := cmd.Run(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

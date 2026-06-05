package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func Contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	builtinCommands := []string{"exit", "echo", "type", "pwd"}

	for {
		fmt.Print("$ ")

		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		parts := strings.Fields(command)
		cmd := parts[0]

		switch cmd {

		case "exit":
			return

		case "echo":
			if len(parts) > 1 {
				fmt.Println(strings.Join(parts[1:], " "))
			} else {
				fmt.Println()
			}

		case "type":
			if len(parts) < 2 {
				continue
			}

			target := parts[1]

			if Contains(builtinCommands, target) {
				fmt.Printf("%s is a shell builtin\n", target)
				continue
			}

			path, err := exec.LookPath(target)
			if err != nil {
				fmt.Printf("%s: not found\n", target)
			} else {
				fmt.Printf("%s is %s\n", target, path)
			}
		case "pwd":
			dir, err := os.Getwd()

			if err != nil{
				log.Fatalf("Failed to read current working directory: %v", err)
			}
			fmt.Println(dir)

		default:
			path, err := exec.LookPath(cmd)

			if err != nil {
				fmt.Printf("%s: command not found\n", cmd)
				continue
			}

			externalCmd := exec.Command(cmd, parts[1:]...)
			externalCmd.Path = path

			externalCmd.Stdout = os.Stdout
			externalCmd.Stderr = os.Stderr

			_ = externalCmd.Run()
		}
	}
}
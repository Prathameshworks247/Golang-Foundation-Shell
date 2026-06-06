package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/google/shlex"
	"os/exec"
	"strings"
)

// Linear search in a slice.
func Contains(slice []string, target string) bool {
	for _, item := range slice { // range returns index, value
		if item == target {
			return true
		}
	}
	return false
}

func main() {
	// Reads input from stdin.
	reader := bufio.NewReader(os.Stdin)

	// Builtin shell commands.
	builtinCommands := []string{"exit", "echo", "type", "pwd", "cd"}

	for {
		fmt.Print("$ ")

		// Read until newline.
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}

		// Remove trailing newline/spaces.
		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		// Split command into tokens.
		parts, err := shlex.Split(command)
		cmd := parts[0]
		// fmt.Print(parts[1:])
		switch cmd {

		case "exit":
			return

		case "echo":
			// Join everything after "echo".
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

			// Check builtins first.
			if Contains(builtinCommands, target) {
				fmt.Printf("%s is a shell builtin\n", target)
				continue
			}

			// Search PATH.
			path, err := exec.LookPath(target)
			if err != nil {
				fmt.Printf("%s: not found\n", target)
			} else {
				fmt.Printf("%s is %s\n", target, path)
			}

		case "pwd":
			// Current working directory.
			dir, err := os.Getwd()
			if err != nil {
				log.Fatalf("Failed to read current working directory: %v", err)
			}
			fmt.Println(dir)

		case "cd":
			if len(parts) < 2 {
				continue
			}

			dir := parts[1]

			// Expand ~ to home directory.
			if dir == "~" {
				dir, _ = os.UserHomeDir()
			}

			// Change shell's working directory.
			err := os.Chdir(dir)
			if err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", parts[1])
			}

		default:
			// Find executable in PATH.
			path, err := exec.LookPath(cmd)
			if err != nil {
				fmt.Printf("%s: command not found\n", cmd)
				continue
			}

			// Create external process.
			externalCmd := exec.Command(cmd, parts[1:]...)

			// Actual executable path.
			externalCmd.Path = path

			// Connect child process output.
			externalCmd.Stdout = os.Stdout
			externalCmd.Stderr = os.Stderr

			// Run and wait.
			_ = externalCmd.Run()
		}
	}
}
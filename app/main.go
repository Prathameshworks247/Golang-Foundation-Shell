package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("$ ")
	
		command, err := reader.ReadString('\n')

		command = strings.TrimSpace(command)

		if command == "exit"{
			break
		}
	
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
	
	
	
		fmt.Println(command + ": command not found")
	}

}

package main

import (
	"fmt"
	"bufio"
	"os"
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
	// TODO: Uncomment the code below to pass the first stage
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("$ ")
	
		command, err := reader.ReadString('\n')
		
		command = strings.TrimSpace(command)
		
		builtin_commands := []string{"exit", "echo", "type"}

		if command == "exit"{
			break
		}else if strings.HasPrefix(command, "echo"){
			fmt.Println(command[5:])
		}else if strings.HasPrefix(command, "type"){
			if Contains(builtin_commands,command[5:]){
				fmt.Println(command[5:] ,"is a shell builtin")
			}else{
				fmt.Println(command[5:] + ": not found")
			}
		}else{
			fmt.Println(command + ": command not found")
		}
	
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
	
	}

}

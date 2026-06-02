package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
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
		cmd := command[5:]
		if command == "exit"{
			break
		}else if strings.HasPrefix(command, "echo"){
			fmt.Println(cmd)
		}else if strings.HasPrefix(command, "type"){
			if Contains(builtin_commands,cmd){ //Imp
				fmt.Println(cmd ,"is a shell builtin")
			}else{

				ans, err := exec.LookPath(cmd)	

				if err == nil {
					fmt.Println(cmd, "is",ans)
					continue
				}

				fmt.Println(cmd + ": not found")
			
			}
		}else{
			fmt.Println(command + ": command not found")
		}
	
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error reading input:", err)//imp
			os.Exit(1)
		}
	
	}

}

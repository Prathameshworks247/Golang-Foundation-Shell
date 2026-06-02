package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	// TODO: Uncomment the code below to pass the first stage
	reader := bufio.NewReader(os.Stdin)
	for {

		fmt.Print("$ ")
	
		command, err := reader.ReadString('\n')
	
		if err != nil{
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
	
	
	
		fmt.Println(command[:len(command)-1] + ": command not found")
	}

}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")

		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = input[:len(input) - 1]
		
		if input == "exit 0" {
			os.Exit(0)
		} else {
			fmt.Fprint(os.Stdout, input + ": command not found\n")
		}
	}
}

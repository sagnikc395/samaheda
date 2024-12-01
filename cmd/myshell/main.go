package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	// fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, "$ ")
		cmd, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		cmd = strings.TrimSpace(cmd)
		cmds := strings.Split(cmd, " ")
		if cmds[0] == "echo" {
			fmt.Println(strings.Join(cmds[1:], " "))
		} else if cmds[0] == "exit" {
			os.Exit(0)
		} else if cmds[0] == "type" {
			if cmds[1] == "echo" {
				fmt.Println("echo is a shell builtin")
			} else if cmds[1] == "exit" {
				fmt.Println("exit is a shell builtin")
			} else if cmds[1] == "type" {
				fmt.Println("type is a shell builtin")
			} else {
				fmt.Println("%s: not found", cmds[1])
			}
		} else {
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

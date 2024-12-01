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

		if cmds[0] == "exit" {
			os.Exit(0)
		}

		fmt.Printf("%s: command not found\n", cmd)
	}
}

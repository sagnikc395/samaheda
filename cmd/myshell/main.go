package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

//check if the program present in PATH
//https://xnacly.me/posts/2023/go-check-for-executable/#:~:text=To%20check%20if%20an%20executable,its%20exported%20LookPath()%20function.

func checkPath(program string) (string, error) {
	pathenv := os.Getenv("PATH")
	paths := strings.Split(pathenv, string(os.PathListSeparator))

	_ = paths
	execpath, err := exec.LookPath(program)
	if err != nil {
		return program, err
	}

	absPath, err := filepath.Abs(execpath)
	if err != nil {
		return program, err
	}

	return absPath, nil
}

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

		if len(cmds) == 0 {
			continue
		}
		switch cmds[0] {
		case "exit":
			os.Exit(0)
		case "echo":
			fmt.Println("echo is a shell builtin")
		case "type":
			if len(cmds) < 2 {
				fmt.Println("type: missing command name")
				continue
			}

			result, err := checkPath(cmds[1])
			if err != nil {
				fmt.Printf("%s: not found\n", cmds[1])
				continue
			}
			fmt.Printf("%s is %s\n", cmds[1], result)
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

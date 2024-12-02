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

// shell builtins
var builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
}

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

func execCommand(program string) error {
	commands := strings.Split(program, " ")
	//first check if present in path or not

	cmdExists, err := checkPath(commands[0])
	if err != nil {
		return err
	}

	cmd := exec.Command(cmdExists, commands[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	_ = out
	fmt.Printf("\n")

	return nil
}

func main() {
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
		case "echo":
			fmt.Println(strings.Join(cmds[1:], " "))
		case "exit":
			os.Exit(0)
		case "type":
			if len(cmds) < 2 {
				fmt.Println("type: missing command name")
				continue
			}

			targetCmd := cmds[1]
			if _, exists := builtins[targetCmd]; exists {
				fmt.Printf("%s is a shell builtin\n", targetCmd)
				continue
			}

			result, err := checkPath(targetCmd)
			if err != nil {
				fmt.Printf("%s: not found\n", targetCmd)
				continue
			}
			fmt.Printf("%s is %s\n", targetCmd, result)

		default:
			// fmt.Printf("%s: command not found\n", cmd)
			err := execCommand(cmd)
			if err != nil {
				fmt.Printf("%s: command not found\n", cmd)
			}
		}
	}
}

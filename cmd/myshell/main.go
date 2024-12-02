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

var builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
}

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

func execCommand(program string) (string, error) {
	commands := strings.Split(program, " ")
	cmdExists, err := checkPath(commands[0])
	if err != nil {
		return "", err
	}

	cmd := exec.Command(cmdExists, commands[1:]...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimRight(string(out), "\r\n"), nil // Trim trailing newlines
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
		if len(cmds) == 0 || cmd == "" {
			continue
		}

		switch cmds[0] {
		case "echo":
			fmt.Print(strings.Join(cmds[1:], " ")) // Remove implicit newline
		case "exit":
			os.Exit(0)
		case "type":
			if len(cmds) < 2 {
				fmt.Print("type: missing command name")
				continue
			}

			targetCmd := cmds[1]
			if _, exists := builtins[targetCmd]; exists {
				fmt.Printf("%s is a shell builtin", targetCmd)
				continue
			}

			result, err := checkPath(targetCmd)
			if err != nil {
				fmt.Printf("%s: not found", targetCmd)
				continue
			}
			fmt.Printf("%s is %s", targetCmd, result)

		default:
			result, err := execCommand(cmd)
			if err != nil {
				fmt.Printf("%s: command not found", cmd)
				continue
			}
			if result != "" {
				fmt.Print(result)
			}
		}
		fmt.Print("\n") // Add a single newline after each command
	}
}

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
	return strings.TrimRight(string(out), "\r\n"), nil
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

		output := ""
		switch cmds[0] {
		case "echo":
			output = strings.Join(cmds[1:], " ")
		case "exit":
			os.Exit(0)
		case "type":
			if len(cmds) < 2 {
				output = "type: missing command name"
				break
			}

			targetCmd := cmds[1]
			if _, exists := builtins[targetCmd]; exists {
				output = fmt.Sprintf("%s is a shell builtin", targetCmd)
				break
			}

			result, err := checkPath(targetCmd)
			if err != nil {
				output = fmt.Sprintf("%s: not found", targetCmd)
				break
			}
			output = fmt.Sprintf("%s is %s", targetCmd, result)

		default:
			result, err := execCommand(cmd)
			if err != nil {
				output = fmt.Sprintf("%s: command not found", cmd)
				break
			}
			output = result
		}

		if output != "" {
			fmt.Print(output)
			fmt.Print("\n")
		}
	}
}

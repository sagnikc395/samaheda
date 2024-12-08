package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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
		case "pwd":
			output = printCurrentWorkingDir()
		case "cd":
			result := changeDirIfExists(cmds[1])
			if !result {
				output = fmt.Sprintf("cd: %s: No such file or directory", cmds[1])
			}
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

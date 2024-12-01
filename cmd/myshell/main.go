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

	// for _, path := range paths {
	// 	if _, err := os.Stat(path); err != nil {
	// 		fmt.Printf("- %s (not accessible)\n", path)
	// 	} else {
	// 		fmt.Printf("- %s\n", path)
	// 	}
	// }

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

		// if cmds[0] == "echo" {
		// 	fmt.Println(strings.Join(cmds[1:], " "))
		// } else if cmds[0] == "exit" {
		// 	os.Exit(0)
		// } else if cmds[0] == "type" {
		// 	result, err := checkPath(cmds[1])
		// 	if err != nil {
		// 		fmt.Printf("%s: not found\n", result)
		// 	}
		// 	fmt.Printf("%s is %s\n", cmds[1], result)
		// } else {
		// 	fmt.Printf("%s: command not found\n", cmd)
		// }

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

			result, err := checkPath(cmds[1])
			if err != nil {
				fmt.Printf("%s: not found\n", cmds[1])
				continue
			}
			if cmds[1] == "echo" {
				fmt.Printf("%s is a shell builtin\n", cmds[1])
			} else {
				fmt.Printf("%s is %s\n", cmds[1], result)
			}
		default:
			fmt.Printf("%s: command not found\n", cmd)
		}
	}
}

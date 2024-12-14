package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var Cmds = make(map[string]func(args []string) error)

func main() {
	Cmds["exit"] = handleExit
	Cmds["echo"] = handleEcho
	Cmds["type"] = handleType
	Cmds["pwd"] = handlePwd
	Cmds["cd"] = handleCd

	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		s := strings.Trim(input, "\r\n")
		var tokens []string
		for {
			start := strings.Index(s, "'")
			if start == -1 {
				tokens = append(tokens, strings.Fields(s)...)
				break
			}
			tokens = append(tokens, strings.Fields(s[:start])...)
			s = s[start+1:]
			end := strings.Index(s, "'")
			token := s[:end]
			tokens = append(tokens, token)
			s = s[end+1:]
		}

		cmd := strings.ToLower(tokens[0])
		var args []string
		if len(tokens) > 1 {
			args = tokens[1:]
		}
		if fn, ok := Cmds[cmd]; ok {
			err := fn(args)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		} else if path, ok := locateCmd(cmd); ok {
			c := exec.Command(path, args...)
			o, err := c.Output()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
			} else {
				fmt.Fprint(os.Stdout, string(o))
			}
		} else {
			fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd)
		}
	}
}

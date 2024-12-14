package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleExit(args []string) error {
	var (
		exitCode int
		err      error
	)

	if len(args) == 1 {
		exitCode, err = strconv.Atoi(args[0])
		if err != nil {
			return err
		}
	}

	os.Exit(exitCode)
	return nil
}

func handleEcho(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(os.Stdout)
		return nil
	}

	for i := 0; i < len(args)-1; i++ {
		fmt.Fprintf(os.Stdout, "%s ", args[i])
	}
	fmt.Fprintln(os.Stdout, args[len(args)-1])
	return nil
}

func locateCmd(cmd string) (string, bool) {
	path := os.Getenv("PATH")
	dirs := strings.Split(path, ":")
	for _, dir := range dirs {
		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		for _, e := range entries {
			if e.IsDir() {
				continue
			}

			parts := strings.Split(e.Name(), ".")
			name := parts[0]
			if cmd == name {
				return fmt.Sprintf("%s/%s", dir, name), true
			}
		}
	}
	return "", false
}

func handleType(args []string) error {
	if len(args) != 1 {
		return nil
	}

	cmd := args[0]
	if _, ok := Cmds[cmd]; ok {
		fmt.Fprintf(os.Stderr, "%s is a shell builtin\n", cmd)
		return nil
	}

	if path, ok := locateCmd(cmd); ok {
		fmt.Fprintf(os.Stdout, "%s is %s\n", cmd, path)
		return nil
	}

	fmt.Fprintf(os.Stderr, "%s: not found\n", cmd)
	return nil
}

func handlePwd(args []string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stdout, dir)
	return nil
}

func handleCd(args []string) error {
	if len(args) == 0 {
		return nil
	}
	dir := args[0]
	if dir == "~" {
		dir = os.Getenv("HOME")
	}
	err := os.Chdir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", dir)
	}
	return nil
}

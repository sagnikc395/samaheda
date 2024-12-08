package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var builtins = map[string]bool{
	"echo": true,
	"exit": true,
	"type": true,
	"pwd":  true,
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

func printCurrentWorkingDir() string {
	dir, err := os.Getwd()
	_ = err
	return dir
}

func changeDirIfExists(path string) bool {

	//check if the path given has ~ prefix , then change to dir
	if strings.HasPrefix(path, "~") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return false
		}

		//replace ~ with actual home dir path
		path = strings.Replace(path, "~", homeDir, 1)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return false
	}

	//check if dir exists
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return false
	}

	//to change to dir
	err = os.Chdir(absPath)
	if err != nil {
		return false
	}
	return true
}

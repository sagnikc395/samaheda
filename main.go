package main

import "fmt"

type Reader struct {
	line   string
	args   *string
	status int
}

func samaheda_loop() {
	for {
		fmt.Printf("> ")
		line, err := samaheda_read_line()
		if err != nil {
			fmt.Printf("Error : %v\n", err)
		}
		args := samahead_split_line(line)
		status := samahead_execute(args)
		if status == EXIT_SUCCESS || status == EXIT_FAILURE {
			break
		}
	}
}

func main() {
	samaheda_loop()
}

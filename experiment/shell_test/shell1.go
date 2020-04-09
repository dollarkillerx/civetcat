package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		cmdString, e := reader.ReadString('\n')
		if e != nil {
			_, e := fmt.Fprintln(os.Stderr, e)
			if e != nil {
				panic(e)
			}

		}
		cmdString = strings.TrimSpace(cmdString)
		command := exec.Command("/bin/bash", "-c", cmdString)
		bufferErr := bytes.NewBufferString("")
		bufferOut := bytes.NewBufferString("")
		command.Stderr = bufio.NewWriter(bufferErr)
		command.Stdout = bufio.NewWriter(bufferOut)
		e = command.Run()
		out := strings.TrimSpace(bufferOut.String())
		errS := strings.TrimSpace(bufferErr.String())
		if out != "" {
			fmt.Println("Out: ")
			fmt.Println(out)
		} else if errS != "" {
			fmt.Println("Err: ")
			fmt.Println(errS)
		}

		if e != nil {
			fmt.Println("Err: ", e)
			continue
		}
	}
}

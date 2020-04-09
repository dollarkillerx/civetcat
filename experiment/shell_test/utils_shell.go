package main

import (
	"bufio"
	"civetcat/utils"
	"fmt"
	"os"
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
		if cmdString == "exit" {
			break
		}

		shell, e := utils.RunShell(cmdString)
		if shell != "" {
			fmt.Println(shell)
		}
	}
}

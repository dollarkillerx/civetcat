/**
 * @Author: DollarKillerX
 * @Description: shell1.go
 * @Github: https://github.com/dollarkillerx
 * @Date: Create in 下午3:48 2019/12/28
 */
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
		cmdString = strings.TrimSuffix(cmdString, "\n")
		command := exec.Command("/bin/bash", "-c", cmdString)
		command.Stderr = os.Stderr
		command.Stdout = os.Stdout
		e = command.Run()
		if e != nil {
			bytes, e := ioutil.ReadAll(os.Stderr)
			if e != nil {
				panic(e)
				if e == io.EOF {
					continue
				}
				log.Println(e)
				continue
			}
			log.Println(string(bytes))
		}
	}
}

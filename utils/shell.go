package utils

import (
	"bytes"
	"os/exec"
	"strings"
)

func RunShell(cmdString string) (string,error) {
	cmdString = strings.TrimSpace(cmdString)
	os := GetOs()
	bufferErr := bytes.NewBufferString("")
	bufferOut := bytes.NewBufferString("")
	var command *exec.Cmd
	switch os {
	case "linux":
		command = exec.Command("/bin/bash", "-c", cmdString)
	default:
		command = exec.Command("cmd", "/C", cmdString)
	}
	command.Stderr = bufferErr
	command.Stdout = bufferOut
	err := command.Run()
	errS := strings.TrimSpace(UTF8(bufferErr.String()))
	outS := strings.TrimSpace(UTF8(bufferOut.String()))
	if err != nil {
		return errS,err
	}
	if outS != "" {
		return outS,nil
	}

	if errS != "" {
		return errS,nil
	}
	return "",nil
}

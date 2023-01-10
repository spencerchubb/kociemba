package kociemba

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func execCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", fmt.Errorf("execCommand StdoutPipe: %v: %v", stdout, err)
	}
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("execCommand Start: %v: %v", stdout, err)
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", fmt.Errorf("execCommand ReadAll: %v: %v", stdout, err)
	}
	cmd.Wait()
	return string(bytes), nil
}

func Solve(facelets string) (string, error) {
	result, err := execCommand("java", "-cp", "./java/dist", "kociemba.Search", facelets)
	if err != nil {
		return "", err
	}
	return result, nil
}

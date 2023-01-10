package kociemba

import (
	"io/ioutil"
	"os/exec"
)

func execCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	if err := cmd.Start(); err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		return "", err
	}
	if err := cmd.Wait(); err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Solve(facelets string) (string, error) {
	result, err := execCommand("java", "-cp", "./java/dist", "kociemba.Search", facelets)
	if err != nil {
		return "", err
	}
	return result, nil
}

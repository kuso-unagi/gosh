package main

import (
	"fmt"
	shell "gosh/internal/shell"
	"io"
	"time"
)

func main() {
	s, err := shell.NewShell("/usr/bin/zsh")
	out := s.GetShellOut()
	go readStdout(out)
	s.Start()
	if err != nil {
		println(err)
	}
	_, err = s.WriteCommand("pwd")
	if err != nil {
		println(err)
	}

	_, err = s.WriteString("ls")
	if err != nil {
		println(err)
	}

	s.Flush()

	_, err = s.WriteCommand("whoami")
	if err != nil {
		println(err)
	}

	time.Sleep(time.Second)
	s.Stop()
	time.Sleep(time.Second)
}

func readStdout(stdout io.Reader) {
	buf := make([]byte, 1)
	for {
		_, err := stdout.Read(buf)
		fmt.Print(string(buf[:]))
		if err == io.EOF {
			println("end")
			break
		}
	}
}

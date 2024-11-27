package shell

import (
	"io"
	"os/exec"
)

type Shell struct {
	shellCmd *exec.Cmd
	shellIn  io.Writer
	shellOut io.Reader
	shellErr io.Reader
}

func NewShell(shellPath string) (*Shell, error) {
	shell := new(Shell)
	shell.shellCmd = exec.Command(shellPath)

	shellIn, err := shell.shellCmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	shellOut, err := shell.shellCmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	shellErr, err := shell.shellCmd.StderrPipe()
	if err != nil {
		return nil, err
	}

	shell.shellIn = shellIn
	shell.shellOut = shellOut
	shell.shellErr = shellErr
	
	return shell, nil
}

func (s *Shell) GetShellIn() io.Writer {
	return s.shellIn
}

func (s *Shell) GetShellOut() io.Reader {
	return s.shellOut
}

func (s *Shell) GetShellErr() io.Reader {
	return s.shellErr
}

func (s *Shell) WriteString(str string) (bool, error) {
	_, err := io.WriteString(s.shellIn, str)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Shell) WriteCommand(str string) (bool, error) {
	str = str + "\n"
	_, err := io.WriteString(s.shellIn, str)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Shell) Flush() (bool, error) {
	_, err := io.WriteString(s.shellIn, "\n")
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *Shell) Start() {
	s.shellCmd.Start()
}

func (s *Shell) Stop() {
	s.shellCmd.Process.Kill()
}

package ido

import (
	"fmt"
	"os/exec"
	"strings"
)

type shell struct {
	line string
	cmd  *exec.Cmd
}

func newShell(name string, arg ...string) *shell {
	s := append([]string{name}, arg[0:]...)
	return &shell{
		line: strings.Join(s, " "),
		cmd:  exec.Command(name, arg...),
	}
}

func (sh *shell) run() error {
	err := sh.cmd.Run()
	if err != nil {
		return fmt.Errorf("%s: %v", sh.line, err)
	}

	return nil
}

func (sh *shell) result() (result string, err error) {
	b, err := sh.cmd.CombinedOutput()
	result = strings.TrimSpace(string(b))
	if err != nil {
		return "", fmt.Errorf("%s: %s", sh.line, result)
	}

	return result, nil
}

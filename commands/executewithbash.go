package commands

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

// ExecuteWithBash -
type ExecuteWithBash struct {
}

// NewExecuteWithBash -
func NewExecuteWithBash() Command {
	return &ExecuteWithBash{}
}

// CanHandle -
func (thisRef *ExecuteWithBash) CanHandle(command string) bool {
	return true
}

// Execute -
func (thisRef *ExecuteWithBash) Execute(command string) error {
	bashParams := strings.Replace(command, "-c", "", 1)

	var customStdErr bytes.Buffer

	cmd := exec.Command("bash", "-c", bashParams)
	cmd.Env = os.Environ()
	cmd.Stdin = os.Stdin
	cmd.Stderr = &customStdErr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return errors.New(customStdErr.String())
	}

	return nil
}

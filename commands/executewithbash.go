package commands

import (
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
func (thisRef *ExecuteWithBash) Execute(command string) {
	bashParams := strings.Replace(command, "-c", "", 1)

	cmd := exec.Command("bash", "-c", bashParams)
	cmd.Env = os.Environ()
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
}

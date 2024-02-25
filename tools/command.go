package tools

import (
	"os"
	"os/exec"
	"strings"
)

// Utility for shell command execution

var shellRuntime = "bash"

type ShellCommand struct {
	Command []string
	Visible bool
}

// Private method for Shell Command struct
func (command *ShellCommand) joinCommand() string {
	return strings.Join(command.Command, " ")
}

func (command *ShellCommand) CommandBuilder() (cmd *exec.Cmd) {
	cmd = exec.Command(shellRuntime, "-c", command.joinCommand())
	cmd.Stdin = os.Stdin
	// If set command as visible state
	if command.Visible {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd
}

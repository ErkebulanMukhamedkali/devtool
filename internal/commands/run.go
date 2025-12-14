package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ErkebulanMukhamedkali/devtool/internal/tools"
)

func Run(command []string) error {
	args := []string{"tool", tools.MODFILE_ARG}

	// go tool -modfile=go.tool.mod govulncheck
	cmd := exec.Command("go", append(args, command...)...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	return nil
}

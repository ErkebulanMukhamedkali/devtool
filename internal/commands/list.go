package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ErkebulanMukhamedkali/devtool/internal/tools"
)

func List() error {
	// go list -modfile=go.tool.mod tool
	cmd := exec.Command("go", "list", tools.MODFILE_ARG, "tool")
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	return nil
}

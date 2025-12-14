package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ErkebulanMukhamedkali/devtool/internal/tools"
)

func Add(name string) error {
	//go get -tool -modfile=go.tool.mod golang.org/x/vuln/cmd/govulncheck
	cmd := exec.Command("go", "get", "-tool", tools.MODFILE_ARG, name)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	return nil
}

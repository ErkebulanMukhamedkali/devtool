package commands

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/ErkebulanMukhamedkali/devtool/internal/tools"
)

func Remove(name string) error {
	//go get -tool -modfile=go.tool.mod golang.org/x/vuln/cmd/govulncheck@none
	cmd := exec.Command("go", "get", "-tool", tools.MODFILE_ARG, fmt.Sprintf("%s@none", name))
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	// comment out for now due to tidy scans all project dependencies to modfile and ironically bloats it

	// cmd = exec.Command("go", "mod", "tidy", tools.MODFILE_ARG)
	// cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	// if err := cmd.Run(); err != nil {
	// 	return fmt.Errorf("failed to tidy: %w", err)
	// }

	return nil
}

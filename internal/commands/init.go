package commands

import (
	"fmt"
	"os"
	"os/exec"

	"devtool/internal/tools"
)

func Init() error {
	name, err := tools.PackageName()
	if err != nil {
		return fmt.Errorf("couldn't get current package name: %w", err)
	}

	// go mod init -modfile=go.tool.mod example.com
	cmd := exec.Command("go", "mod", "init", tools.MODFILE_ARG, fmt.Sprintf("%s-tools", name))
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("failed to run command: %w", err)
	}

	return nil
}

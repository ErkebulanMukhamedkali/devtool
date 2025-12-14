package tools

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	GO_MOD_FILE = "go.mod"
	MODFILE_ARG = "-modfile=go.tool.mod"
)

var ErrIncorrectFormat = errors.New("invalid mod file")

func PackageName() (string, error) {
	f, err := os.Open(GO_MOD_FILE)
	if err != nil {
		return "", fmt.Errorf("error reading go.mod on path %s: %w", GO_MOD_FILE, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if firstLine := strings.Fields(scanner.Text()); len(firstLine) == 2 {
			return firstLine[1], nil
		} else {
			return "", ErrIncorrectFormat
		}
	}

	return "", fmt.Errorf("something went wrong: %w", scanner.Err())
}

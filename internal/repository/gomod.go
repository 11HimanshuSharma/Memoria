package repository

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// findGoMod walks upwards looking for go.mod
func findGoMod(start string) (string, string, bool, error) {
	current, err := filepath.Abs(start)
	if err != nil {
		return "", "", false, err
	}
	for {
		modPath := filepath.Join(current, "go.mod")

		if _, err := os.Stat(modPath); err == nil {
			module, err := parseModule(modPath)

			if err != nil {
				return "", "", false, err
			}
			return current, module, true, nil
		}
		parent  := filepath.Dir(current)

		if parent == current {
			return "", "", false, nil
		}
		current = parent
	}
}



func parseModule(path string) (string, error){
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "module "){
			return strings.TrimSpace(strings.TrimPrefix(line, "module ")), nil
		}
	}
	return "", scanner.Err()
}
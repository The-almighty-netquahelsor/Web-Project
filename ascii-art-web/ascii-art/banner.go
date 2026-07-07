package asciiart

import (
	"fmt"
	"os"
	"strings"
)

const requiredLines = 95 * 9

func LoadBanner(name string) ([]string, error) {
	filename := strings.TrimSuffix(name, ".txt") + ".txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("banner %q not found", name)
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("banner file %q is empty", filename)
	}
	lines := strings.Split(string(data), "\n")

	if len(lines) < requiredLines {
		return nil, fmt.Errorf("banner file %q is incomplete: expected at least %d lines, got %d", filename, requiredLines, len(lines))
	}
	return lines, nil
}

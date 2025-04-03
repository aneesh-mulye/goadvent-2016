package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func DayFileLines(day uint) ([]string, error) {
	f, err := os.Open(filepath.Join("inputs", fmt.Sprintf("day%02d.txt", day)))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

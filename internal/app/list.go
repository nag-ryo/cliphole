package app

import (
	"bufio"
	"os"
	"path/filepath"
)

func LoadHistoryLines() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(home, ".cliphole", "history.txt")

	file, err := os.Open(path)
	if err != nil {
		// ファイルが存在しない場合は空で返す
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, scanner.Err()
}

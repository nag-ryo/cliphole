package history

import (
	"os"
	"path/filepath"
)

type FileStore struct {
	Path string
}

func NewFileStore() (*FileStore, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(home, ".cliphole", "history.txt")
	return &FileStore{Path: path}, nil
}

func (f *FileStore) Append(text string) error {
	if err := os.MkdirAll(filepath.Dir(f.Path), 0755); err != nil {
		return err
	}
	file, err := os.OpenFile(f.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(text + "\n")
	return err
}

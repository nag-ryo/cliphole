package app

import "github.com/yourname/cliphole/internal/history"

func SaveText(text string) error {
	store, err := history.NewFileStore()
	if err != nil {
		return err
	}
	return store.Append(text)
}

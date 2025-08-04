package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [text]",
	Short: "Save a string to clipboard history",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := args[0]
		historyPath, err := getHistoryFilePath()
		if err != nil {
			return err
		}

		// ディレクトリがなければ作成
		if err := os.MkdirAll(filepath.Dir(historyPath), 0755); err != nil {
			return err
		}

		// ファイルに追記
		f, err := os.OpenFile(historyPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.WriteString(text + "\n"); err != nil {
			return err
		}

		fmt.Println("Saved to history.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}

// ~/.cliphole/history.txt のフルパスを返す
func getHistoryFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".cliphole", "history.txt"), nil
}

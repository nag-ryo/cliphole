package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/yourname/cliphole/internal/app"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Select from history and copy to clipboard",
	RunE: func(cmd *cobra.Command, args []string) error {
		items, err := app.LoadHistoryLines()
		if err != nil {
			return err
		}
		if len(items) == 0 {
			fmt.Println("No history found.")
			return nil
		}

		// fzf 実行
		selected, err := runFzf(items)
		if err != nil {
			return err
		}
		if selected == "" {
			fmt.Println("Nothing selected.")
			return nil
		}

		// pbcopy へ送信
		if err := writeToClipboard(selected); err != nil {
			return err
		}

		fmt.Println("Copied to clipboard.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

// fzf を使って選択させる
func runFzf(lines []string) (string, error) {
	cmd := exec.Command("fzf")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", err
	}
	go func() {
		defer stdin.Close()
		io := strings.Join(lines, "\n")
		stdin.Write([]byte(io))
	}()

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// pbcopy に送る
func writeToClipboard(text string) error {
	cmd := exec.Command("pbcopy")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := cmd.Start(); err != nil {
		return err
	}

	_, err = stdin.Write([]byte(text))
	if err != nil {
		return err
	}

	stdin.Close()
	return cmd.Wait()
}

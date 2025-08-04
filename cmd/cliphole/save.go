package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yourname/cliphole/internal/app"
)

var saveCmd = &cobra.Command{
	Use:   "save [text]",
	Short: "Save a string to clipboard history",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		text := args[0]
		if err := app.SaveText(text); err != nil {
			return err
		}
		fmt.Println("Saved to history.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)
}

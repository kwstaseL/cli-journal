/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

const defaultLimit = 5;

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List recently added notes",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		if limit < 0 {
			limit = defaultLimit
		}
		frequentNotes, err := noteService.ListFrequentNotes(limit)
		if err != nil {
			noteCLIDrawer.DrawError("Error while trying to find recently added notes", err)
			return
		}
		noteCLIDrawer.DrawListNotes(frequentNotes)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntP("limit", "l", defaultLimit, "Number of notes to show")
}

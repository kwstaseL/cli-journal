/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List recently added notes",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		notes, _ := noteService.ListFrequentNotes(limit)
		fmt.Println(notes)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntP("limit", "l", 3, "Amount of notes to show")
}

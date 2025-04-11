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
		frequentNotes, err := noteService.ListFrequentNotes(limit)
		if err != nil {
			fmt.Printf("Error: Unable to find recent notes. %v\n", err)
			return
		}
		// TODO: Send a proper message to the user
		fmt.Println(frequentNotes)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntP("limit", "l", 3, "Number of notes to show")
}

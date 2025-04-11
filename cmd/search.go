/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/kwstaseL/cli-journal/pkg/model"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for your notes by tag, category or text.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		tags, _ := cmd.Flags().GetString("tags")
		category, _ := cmd.Flags().GetString("category")
		text, _ := cmd.Flags().GetString("text")
		filteredNotes, _ := noteService.ListNotesBy(model.NoteFilters{Tags: tags, Category: category, SearchTerm: text})
		fmt.Println(filteredNotes)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("tags", "t", "", "Tags to look for (comma-separated)")
	searchCmd.Flags().StringP("category", "c", "", "Category to look for")
	searchCmd.Flags().StringP("text", "b", "", "Text to look for")
}

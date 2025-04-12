/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"github.com/kwstaseL/cli-journal/pkg/model"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for your notes by tag, category or text.",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		searchTags, _ := cmd.Flags().GetString("tags")
		searchCategory, _ := cmd.Flags().GetString("category")
		searchTerm, _ := cmd.Flags().GetString("text")
		filteredNotes, err := noteService.ListNotesBy(model.NoteFilters{Tags: searchTags, Category: searchCategory,
																		 SearchTerm: searchTerm})
		if err != nil {
			noteCLIDrawer.DrawError("Error while trying to search for notes", err)
			return
		}
		noteCLIDrawer.DrawListNotes(filteredNotes)
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("tags", "t", "", "Tags to look for (comma-separated)")
	searchCmd.Flags().StringP("category", "c", "", "Category to look for")
	searchCmd.Flags().StringP("text", "b", "", "Text to look for")
}

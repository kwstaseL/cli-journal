package cmd

import (
	"github.com/kwstaseL/cli-journal/pkg/model"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a new note",
	Long:  "Save a new note by header, body, category, and tags",
	Run: func(cmd *cobra.Command, args []string) {
		header, _ := cmd.Flags().GetString("header")
		body, _ := cmd.Flags().GetString("body")
		category, _ := cmd.Flags().GetString("category")
		tags, _ := cmd.Flags().GetString("tags")
		
		note := model.NewNote(header, body, category, tags)
		err := noteService.CreateNewNote(*note)
		
		if err != nil {
			noteCLIDrawer.DrawError("Error while trying to save note", err)
			return
		}
		noteCLIDrawer.DrawSaveSuccess(*note)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.Flags().StringP("header", "t", "", "Header of the note")
	saveCmd.Flags().StringP("body", "b", "", "Body of the note")
	saveCmd.Flags().StringP("category", "c", "general", "Category of the note")
	saveCmd.Flags().StringP("tags", "g", "", "Tags for the note (comma-separated)")

	saveCmd.MarkFlagRequired("header")
	saveCmd.MarkFlagRequired("body")
}
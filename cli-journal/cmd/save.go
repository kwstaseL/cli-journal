/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/kwstaseL/cli-journal/cmd/internal/model"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a new note",
	Long:  "Save a new note with title, content, category, and tags",
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		content, _ := cmd.Flags().GetString("content")
		category, _ := cmd.Flags().GetString("category")
		tags, _ := cmd.Flags().GetStringSlice("tags")

		fmt.Printf("Saving note with title: %s\n", title)
		fmt.Printf("Content: %s\n", content)
		fmt.Printf("Category: %s\n", category)
		fmt.Printf("Tags: %v\n", tags)


		note := model.NewNote(title, content, category, tags)
		// TODO: Implement actual save functionality
		fmt.Println("Note saved successfully!", note)
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.Flags().StringP("title", "t", "", "title of the note")
	saveCmd.Flags().StringP("content", "c", "", "content of the note")
	saveCmd.Flags().StringP("category", "g", "general", "category of the note")
	saveCmd.Flags().StringSliceP("tags", "T", []string{}, "tags for the note (comma-separated)")

	saveCmd.MarkFlagRequired("title")
	saveCmd.MarkFlagRequired("content")
}
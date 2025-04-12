/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"os"

	logger "github.com/kwstaseL/cli-journal/pkg/logging"
	"github.com/kwstaseL/cli-journal/pkg/service"
	"github.com/spf13/cobra"
)

var noteService service.NoteService
var noteCLIDrawer *NoteCLIDrawer

func SetNoteService(service service.NoteService) {
	noteService = service
}

var rootCmd = &cobra.Command{
	Use:   "cli-journal",
	Short: "A simple CLI tool to save, edit and share notes",
	Long: ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		logger.LogError("Error occured while trying to execute root.")
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	noteCLIDrawer = NewNoteCLIDrawer()
}

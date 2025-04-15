/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"github.com/kwstaseL/cli-journal/pkg/export"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a note to a file format like txt, markdown, or json",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		platform, _ := cmd.Flags().GetString("format")

		note, err := noteService.GetNoteById(id)
		if err != nil {
			noteCLIDrawer.DrawError("Error while trying to find note", err)
			return
		}

		exporter, err := export.GetExporter(platform)
		if err != nil {
			noteCLIDrawer.DrawError("Error while trying to export note ", err)
			return
		}

		err = exporter.Export(*note)
		if err != nil {
			noteCLIDrawer.DrawError("Failed to export note", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(exportCmd)
	exportCmd.Flags().StringP("id", "i", "", "ID of the note to export")
	exportCmd.Flags().StringP("format", "f", "txt", "Format to export the note as txt, json or md")

	exportCmd.MarkFlagRequired("id")
	exportCmd.MarkFlagRequired("format")
}
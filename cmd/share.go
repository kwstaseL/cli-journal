/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"github.com/kwstaseL/cli-journal/pkg/share"
	"github.com/spf13/cobra"
)

var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "Share your note to a platform like X or Email",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetString("id")
		platform, _ := cmd.Flags().GetString("platform")

		note, err := noteService.GetNoteById(id)
		if err != nil {
			noteCLIDrawer.DrawError("Failed to find note", err)
			return
		}

		sharer, err := share.GetSharer(platform)
		if err != nil {
			noteCLIDrawer.DrawError("Invalid share platform", err)
			return
		}

		err = sharer.Share(*note)
		if err != nil {
			noteCLIDrawer.DrawError("Failed to share note", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)

	shareCmd.Flags().StringP("platform", "p", "", "Platform to share the note (e.g., x, email)")
	shareCmd.Flags().StringP("id", "i", "", "ID of the note to share")

	shareCmd.MarkFlagRequired("platform")
	shareCmd.MarkFlagRequired("id")
}

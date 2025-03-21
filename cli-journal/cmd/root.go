/*
Copyright © 2025 Konstantinos Eleftheriou <eleftheriou.konst@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var dbConnection *gorm.DB

var rootCmd = &cobra.Command{
	Use:   "cli-journal",
	Short: "A simple CLI tool to save, edit and share notes",
	Long: `A simple CLI tool to save, edit and share notes`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

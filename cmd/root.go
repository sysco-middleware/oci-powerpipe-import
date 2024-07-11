/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "importer",
	Short: "Import powerpipe-reports to database",
	Long: `Import given powerpipe reports to streampipe database.
Currently only supports import of oci compliance report in csv format.

Flags marked as * are required

`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.PersistentFlags().StringP(dbUser, dbUserShort, "steampipe", "* Database user")
	rootCmd.PersistentFlags().StringP(dbHost, dbHostShort, "127.0.0.1", "* Database host")
	rootCmd.PersistentFlags().StringP(dbPass, dbPassShort, "", "* Database password")
	rootCmd.PersistentFlags().IntP(dbPort, dbPortShort, 9133, "* Database port")
	rootCmd.PersistentFlags().StringP(dbService, dbServiceShort, "steampipe", "* Database service")
	rootCmd.PersistentFlags().StringP(importFilePath, importFilePathShort, "", "* Import file path")

	_ = rootCmd.MarkPersistentFlagRequired(importFilePath)
	_ = rootCmd.MarkPersistentFlagRequired(dbUser)
	_ = rootCmd.MarkPersistentFlagRequired(dbHost)
	_ = rootCmd.MarkPersistentFlagRequired(dbPass)
	_ = rootCmd.MarkPersistentFlagRequired(dbPort)
	_ = rootCmd.MarkPersistentFlagRequired(dbService)
}

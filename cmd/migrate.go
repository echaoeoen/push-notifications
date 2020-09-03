package cmd

import (
	"github.com/oeoen/push-notifications/cmd/cli"
	"github.com/spf13/cobra"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Various migration helpers",
	Run:   cli.MigrateSQLHandler(NewDriver()),
}

func init() {
	RootCmd.AddCommand(migrateCmd)
}

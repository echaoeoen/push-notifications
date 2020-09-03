package cli

import (
	"github.com/oeoen/push-notifications/driver"
	"github.com/oeoen/push-notifications/pkg/storage/sqls"
	"github.com/spf13/cobra"
)

func MigrateSQLHandler(d driver.Driver) func(*cobra.Command, []string) {
	return func(cmd *cobra.Command, args []string) {
		sqlsDB := d.Registry().NotificationManager().StorageManager().(*sqls.SQLs)
		sqlsDB.DoMigration(d.Configuration(), args...)
	}
}

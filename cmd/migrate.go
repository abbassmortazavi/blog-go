package cmd

import (
	"blog/pkg/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(migrateCMD)
}

var migrateCMD = &cobra.Command{
	Use:   "migrate",
	Short: "Table Migration",
	Long:  "Table Migration",
	Run: func(cmd *cobra.Command, args []string) {
		migrate()
	},
}

func migrate() {
	bootstrap.Migrate()
}

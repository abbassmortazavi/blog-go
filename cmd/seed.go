package cmd

import (
	"blog/pkg/bootstrap"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(seedCMD)
}

var seedCMD = &cobra.Command{
	Use:   "seed",
	Short: "Table Seeder",
	Long:  "Table Seeder",
	Run: func(cmd *cobra.Command, args []string) {
		seed()
	},
}

func seed() {
	bootstrap.Seed()
}

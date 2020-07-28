package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"todo/app"
	"todo/model"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {
		refresh, _ := cmd.Flags().GetBool("refresh")

		a, err := app.New("")
		if err != nil {
			return err
		}
		defer a.Database.CloseDB()

		if refresh {
			fmt.Println("dropping tables...")
			a.Database.DropTables(
				&model.Todo{})
		}

		a.Database.Migrate(
			&model.Todo{})
		fmt.Println("migration success!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().Bool("refresh", false, "drop all tables before migration")
}

package main

import (
	"github.com/frain-dev/convoy/config"
	"github.com/frain-dev/convoy/database"
	"github.com/frain-dev/convoy/internal/pkg/migrator"
	"github.com/frain-dev/convoy/pkg/log"

	"github.com/spf13/cobra"
)

func addMigrateCommand(a *app) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "Convoy migrations",
	}

	cmd.AddCommand(addUpCommand())
	cmd.AddCommand(addDownCommand())

	return cmd
}

func addUpCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "up",
		Aliases: []string{"migrate-up"},
		Short:   "Run all pending migrations",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := config.Get()
			if err != nil {
				log.WithError(err).Fatalf("Error fetching the config.")
			}

			db, err := database.New()
			if err != nil {
				log.Fatal(err)
			}

			m := migrator.New(db)
			err = m.Up()
			if err != nil {
				log.Fatalf("migration up failed with error: %+v", err)
			}
		},
	}

	return cmd
}

func addDownCommand() *cobra.Command {
	var migrationID string

	cmd := &cobra.Command{
		Use:     "down",
		Aliases: []string{"migrate-down"},
		Short:   "Rollback migrations",
		Run: func(cmd *cobra.Command, args []string) {
			_, err := config.Get()
			if err != nil {
				log.WithError(err).Fatalf("Error fetching the config.")
			}

			db, err := database.New()
			if err != nil {
				log.Fatal(err)
			}

			m := migrator.New(db)
			err = m.Down()
			if err != nil {
				log.Fatalf("migration up failed with error: %+v", err)
			}
		},
	}

	cmd.Flags().StringVar(&migrationID, "id", "", "Migration ID")

	return cmd
}

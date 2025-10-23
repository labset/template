package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "clarity",
	}

	rootCmd.AddCommand(DbSchemaRollbackCmd())
	rootCmd.AddCommand(DbSchemaMigrateCmd())

	err := rootCmd.Execute()
	if err != nil {
		log.Fatalf("failed to execute command: %v", err)
	}
}

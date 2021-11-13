package main

import (
	"context"
	"log"

	"github.com/lucasres/db-stack/cmd"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dbstack",
		Short: "DBstack is a key-value database in memory but like stack",
		Long:  `This is a key-value database  in memory wherekey are stack what you can get current value and pop for previous version of the key`,
		Run: func(c *cobra.Command, args []string) {
			cmd.HandleVersionCmd()
		},
	}

	rootCmd.AddCommand(cmd.NewVersionCmd())

	rootCtx := context.Background()

	if err := rootCmd.ExecuteContext(rootCtx); err != nil {
		log.Fatal(err)
	}
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show current version",
		Run: func(cmd *cobra.Command, args []string) {
			HandleVersionCmd()
		},
	}
}

func HandleVersionCmd() {
	fmt.Println("Version 0.1 👽")
}

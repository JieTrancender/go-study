package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genVersionCmd generates the version command
func genVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Show current version info",
		Run: func(_ *cobra.Command, args []string) {
			fmt.Println("kbm v0.0.1")
		},
	}
}

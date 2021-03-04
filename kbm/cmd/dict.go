package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func genDictCmd() *cobra.Command {
	word := ""

	dictCmd := &cobra.Command{
		Use:   "dict",
		Short: "kbm dict tool",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("kbm dict word(%s)\n", word)
		},
	}

	dictCmd.Flags().StringVarP(&word, "word", "w", "hello", "search word")

	return dictCmd
}

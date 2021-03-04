package cmd

import (
	"github.com/spf13/cobra"
)

// KbmRootCmd indictes cmd structure
type KbmRootCmd struct {
	cobra.Command
	RunCmd     *cobra.Command
	VersionCmd *cobra.Command
	DictCmd    *cobra.Command
}

// NewRootCommand creates root command
func NewRootCommand() *KbmRootCmd {
	rootCmd := &KbmRootCmd{}

	rootCmd.RunCmd = genRunCommand()
	rootCmd.VersionCmd = genVersionCmd()
	rootCmd.DictCmd = genDictCmd()

	rootCmd.Run = rootCmd.RunCmd.Run

	rootCmd.AddCommand(rootCmd.RunCmd)
	rootCmd.AddCommand(rootCmd.VersionCmd)
	rootCmd.AddCommand(rootCmd.DictCmd)

	return rootCmd
}

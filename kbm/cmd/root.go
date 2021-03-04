package cmd

import (
	// "fmt"
	// "strings"

	"github.com/spf13/cobra"
	// "github.com/spf13/pflag"
	// "github.com/spf13/viper"
)

// KbmRootCmd indictes cmd structure
type KbmRootCmd struct {
	cobra.Command
	RunCmd     *cobra.Command
	VersionCmd *cobra.Command
}

// NewRootCommand creates root command
func NewRootCommand() *KbmRootCmd {
	rootCmd := &KbmRootCmd{}

	rootCmd.RunCmd = genRunCommand()
	rootCmd.VersionCmd = genVersionCmd()

	rootCmd.Run = rootCmd.RunCmd.Run

	rootCmd.AddCommand(rootCmd.RunCmd)
	rootCmd.AddCommand(rootCmd.VersionCmd)

	return rootCmd
}

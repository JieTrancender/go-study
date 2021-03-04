package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	// DefaultConfigFilename indicates the default config file name.
	DefaultConfigFilename = "kbm"

	// EnvPrefix sets the prefix of Env
	EnvPrefix = "KBM"
)

// NewRootCommand creates root command
func NewRootCommand() *cobra.Command {
	etcdHost := []string{}
	etcdPath := ""

	rootCmd := &cobra.Command{
		Use:   "kbm",
		Short: "kbm tools",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return initializeConfig(cmd)
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("kbm etcd-host(%v) etcd-path(%s)\n", etcdHost, etcdPath)
		},
	}

	rootCmd.Flags().StringArrayVar(&etcdHost, "etcd-host", []string{"127.0.0.1:2379"}, "etcd host(may be given multi times)")
	rootCmd.Flags().StringVar(&etcdPath, "etcd-path", "/config/kbm/default", "path of the config in etcd")

	return rootCmd
}

func initializeConfig(cmd *cobra.Command) error {
	v := viper.New()

	v.SetConfigName(DefaultConfigFilename)

	v.AddConfigPath(".")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	v.SetEnvPrefix(EnvPrefix)

	v.AutomaticEnv()

	bindFlags(cmd, v)

	return nil
}

func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			envVarSuffix := strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_"))
			v.BindEnv(f.Name, fmt.Sprintf("%s_%s", EnvPrefix, envVarSuffix))
		}

		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}

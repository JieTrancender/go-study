package main

import (
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	if *cfg != "" {
		viper.SetConfigFile(*cfg)
		viper.SetConfigType("yaml")
	} else {
		viper.AddConfigPath("./conf")
		viper.AddConfigPath("$HOME/.excel_to_lua")
		viper.SetConfigName("config")
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	// viper.WatchConfig()
	// viper.OnConfigChange(func(e fsnotify.Event) {
	// 	fmt.Println("Config file changed:", e.Name)
	// })

	fmt.Printf("Used configuration file is: %s\n", viper.ConfigFileUsed())

	// var version bool
	// flagSet := pflag.NewFlagSet("test", pflag.ContinueOnError)
	// flagSet.BoolVar(&version, "version", true, "Print verison information and quit.")

	// pflag.Parse()

	// dir := "./excelDir/"
	// files, err := ioutil.ReadDir(dir)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for _, file := range files {
	// 	fmt.Println(file.Name())
	// }
}

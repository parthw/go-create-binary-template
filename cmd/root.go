package cmd

import (
	"fmt"
	"os"

	"github.com/parthw/go-create-binary-template/internal/hello"
	"github.com/parthw/go-create-binary-template/internal/logger"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-create-binary-template",
	Short: "A brief description of your application",
	Run: func(cmd *cobra.Command, args []string) {
		hello.Hello()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-create-binary-template.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// Setting the default configuration
	defaultConfig()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".go-create-binary-template" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-create-binary-template")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	// Initializing Logger of application
	logger.InitializeLogger()
}

func defaultConfig() {
	viper.SetDefault("log.type", "console")
	viper.SetDefault("log.file", "go-create-binary-template.log")
	viper.SetDefault("log.file.maxsize", "100") //megabytes
	viper.SetDefault("log.file.maxbackups", "5")
	viper.SetDefault("log.file.maxage", "5") //days
	viper.SetDefault("log.level", "info")
}

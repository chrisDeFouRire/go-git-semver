package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
	force     bool
	assumeYes bool
	message   string

	meta       string
	prerelease string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-semver",
	Short: "Manage Semver tags easily in git",
	Long:  `Manage Semver tags easily in git`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string) {
	rootCmd.Version = version

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-git-semver.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&force, "force", "f", false, "force tag, don't warn if dirty folder or already tagged")
	rootCmd.PersistentFlags().BoolVarP(&assumeYes, "yes", "y", false, "Assume Yes answer")
	rootCmd.PersistentFlags().StringVarP(&message, "msg", "m", "", "message for annotated tag")

	rootCmd.PersistentFlags().StringVar(&prerelease, "prerelease", "", "specify a prerelease suffix")
	rootCmd.PersistentFlags().StringVar(&meta, "meta", "", "specify a metadata suffix")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".go-git-semver" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".go-git-semver")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

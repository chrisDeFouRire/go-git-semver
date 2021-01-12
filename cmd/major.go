package cmd

import (
	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"
)

// majorCmd represents the major command
var majorCmd = &cobra.Command{
	Use:   "major",
	Short: "Bump the major level",
	Long:  `Bump the major level`,

	Run: bumpRepoWithBumper(semver.Version.IncMajor),
}

func init() {
	rootCmd.AddCommand(majorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// majorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// majorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

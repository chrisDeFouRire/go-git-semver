package cmd

import (
	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"
)

// minorCmd represents the minor command
var minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Bump the minor level",
	Long:  `Bump the minor level`,

	Run: bumpRepoWithBumper(semver.Version.IncMinor),
}

func init() {
	rootCmd.AddCommand(minorCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// minorCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// minorCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

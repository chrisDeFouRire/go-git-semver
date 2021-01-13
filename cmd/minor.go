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
}

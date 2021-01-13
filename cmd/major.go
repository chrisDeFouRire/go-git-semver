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

}

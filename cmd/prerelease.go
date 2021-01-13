package cmd

import (
	"log"

	"github.com/Masterminds/semver"
	"github.com/spf13/cobra"
)

var meta string

// prereleaseCmd represents the prerelease command
var prereleaseCmd = &cobra.Command{
	Use:   "prerelease [--meta=metadata] pre-release",
	Short: "prerelease lets you tag with a pre-release semver suffix",
	Long:  `prerelease lets you tag with a pre-release semver suffix`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		bumper := func(in semver.Version) semver.Version {
			tmpv, err := in.SetPrerelease(args[0])
			if err != nil {
				log.Fatal(err)
			}
			if meta != "" {
				tmpv, err = in.SetMetadata(meta)
				if err != nil {
					log.Fatal(err)
				}
			}
			return tmpv
		}
		bumpRepoWithBumper(bumper)(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(prereleaseCmd)

	prereleaseCmd.Flags().StringVar(&meta, "meta", "", "specify a metadata suffix")
}

package cmd

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"

	"github.com/chrisDeFouRire/gitv/lib"
)

var nov, nonl bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the latest semver tag",
	Long:  `Get the latest semver tag`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}

		tag, _, _, err := lib.FindLatestSemverTag(repo)
		if err != nil {
			log.Fatal(err)
		}
		if !nonl {
			defer fmt.Println()
		}

		if nov {
			fmt.Print(tag[1:])
			return
		}
		fmt.Print(tag)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().BoolVarP(&nov, "nov", "v", false, "Removes the prefix v letter")
	getCmd.Flags().BoolVarP(&nonl, "nonl", "n", false, "Removes trailing newline character")
}

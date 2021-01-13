package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

var origin string

// purgeCmd represents the purge command
var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Removes tags according to semver constraint",
	Long:  `Removes tags according to semver constraint`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		c, err := semver.NewConstraint(args[0])
		if err != nil {
			log.Fatal(err)
		}

		repo, err := git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}
		tags, err := repo.Tags()
		if err != nil {
			log.Fatal(err)
		}

		var tagList []string
		for ref, err := tags.Next(); err == nil; ref, err = tags.Next() {
			tagName := ref.Name().Short()
			if version, notVersionErr := semver.NewVersion(tagName); notVersionErr == nil {
				if c.Check(version) {
					tagList = append(tagList, tagName)
				}
			}
		}
		tags.Close()

		if len(tagList) == 0 {
			fmt.Println("No tags match", args[0])
			os.Exit(0)
		}

		fmt.Println("Tags that match", args[0])
		for _, tag := range tagList {
			fmt.Println(tag)
		}
		fmt.Print("Purge? y/N: ")
		var ok string
		fmt.Fscan(os.Stdin, &ok)
		if ok == "y" {
			for _, tag := range tagList {
				cmd := exec.Command("git", "push", origin, ":refs/tags/"+tag)
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Removed", tag)

				err = repo.DeleteTag(tag)
				if err != nil {
					log.Fatal(err)
				}
			}
			return
		}
		fmt.Println("Purge canceled")
	},
}

func init() {
	purgeCmd.Flags().StringVar(&origin, "origin", "origin", "specify the origin where tag deletion should be pushed")
	rootCmd.AddCommand(purgeCmd)
}

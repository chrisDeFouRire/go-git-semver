package cmd

import (
	"fmt"
	"log"

	"github.com/chrisDeFouRire/gitv/lib"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

// patchCmd represents the patch command
var patchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Bump the patch level",
	Long:  `Bump the patch level`,
	Run: func(cmd *cobra.Command, args []string) {
		repo, err := git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}

		dirty, err := lib.DirtyFolder(repo)
		if err != nil {
			log.Fatal(err)
		}
		if dirty {
			log.Fatal("Directory is dirty, commit first")
		}

		tag, hash, v, err := lib.FindLatestSemverTag(repo)
		if err != nil {
			log.Fatal(err)
		}
		if !nonl {
			defer fmt.Println()
		}

		head, err := repo.Head()
		if hash.String() == head.Hash().String() {
			log.Fatalf("No need to bump, tag %s applies to HEAD", tag)
		}

		newVersion := v.IncPatch()
		newTag := "v" + newVersion.String()
		log.Printf("Tagging with tag %s", newTag)

	},
}

func init() {
	rootCmd.AddCommand(patchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// patchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// patchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

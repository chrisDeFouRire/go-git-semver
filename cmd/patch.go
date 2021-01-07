package cmd

import (
	"fmt"
	"log"
	"os"

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
		if dirty && !quiet {
			log.Fatal("Directory is dirty, commit first")
		}

		tag, hash, v, err := lib.FindLatestSemverTag(repo)
		if err != nil {
			log.Fatal(err)
		}

		head, err := repo.Head()
		if hash.String() == head.Hash().String() {
			log.Fatalf("No need to bump, tag %s applies to HEAD", tag)
		}

		newVersion := v.IncPatch()
		newTag := "v" + newVersion.String()
		if nov {
			newTag = newVersion.String()
		}

		ok := "n"
		if !assumeYes {
			fmt.Printf("Tag with %s? (y/N)  ", newTag)
			fmt.Fscan(os.Stdin, &ok)
		}
		if assumeYes || ok == "y" {
			ref, err := repo.CreateTag(newTag, head.Hash(), nil) // nil to create non annotated tag
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Tagged %s with tag %s\n", ref.String(), newTag)
		} else {
			if !quiet {
				fmt.Println("Exit without tagging")
			}
			os.Exit(-1)
		}
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

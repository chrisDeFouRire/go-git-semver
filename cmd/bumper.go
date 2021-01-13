package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Masterminds/semver"
	"github.com/chrisDeFouRire/go-git-semver/lib"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
)

type bumper func(v semver.Version) semver.Version

func bumpRepoWithBumper(bump bumper) func(*cobra.Command, []string) {

	return func(cmd *cobra.Command, args []string) {
		repo, err := git.PlainOpen(".")
		if err != nil {
			log.Fatal(err)
		}

		dirty, err := lib.DirtyFolder(repo)
		if err != nil {
			log.Fatal(err)
		}
		if dirty && !force {
			log.Fatal("Directory is dirty, commit first")
		}

		tag, hash, v, err := lib.FindLatestSemverTag(repo)
		if err != nil {
			log.Fatal(err)
		}

		head, err := repo.Head()
		if hash.String() == head.Hash().String() && !force {
			log.Fatalf("No need to bump, tag %s applies to HEAD", tag)
		}

		if clear {
			var tmpv semver.Version
			if tmpv, err = v.SetPrerelease(""); err != nil {
				log.Fatal(err)
			}
			if tmpv, err = v.SetMetadata(""); err != nil {
				log.Fatal(err)
			}
			v = &tmpv
		}
		newVersion := bump(*v)
		if prerelease != "" {
			newVersion, err = newVersion.SetPrerelease(args[0])
			if err != nil {
				log.Fatal(err)
			}
		}
		if meta != "" {
			newVersion, err = newVersion.SetMetadata(meta)
			if err != nil {
				log.Fatal(err)
			}
		}

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
			var opts *git.CreateTagOptions
			if message != "" {
				opts = &git.CreateTagOptions{Message: message}
			}
			ref, err := repo.CreateTag(newTag, head.Hash(), opts) // nil to create non annotated tag
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Tagged %s with tag %s\n", ref.String(), newTag)
		} else {
			if !force {
				fmt.Println("Exit without tagging")
			}
			os.Exit(-1)
		}
	}
}

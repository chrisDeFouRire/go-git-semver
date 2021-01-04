package main

import (
	"fmt"
	"log"

	"github.com/Masterminds/semver"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {

	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal(err)
	}

	wt, err := repo.Worktree()
	if err != nil {
		log.Fatal(err)
	}
	status, err := wt.Status()
	if err != nil {
		log.Fatal(err)
	}
	if status.IsClean() {
		log.Println("Directory is clean")
	} else {
		log.Fatal("Directory is dirty")
	}

	tag, htag, err := FindLatestSemverTag(repo)
	fmt.Println("Latest semver tag found on current branch: ", tag)

	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	if htag.String() == head.Hash().String() {
		log.Fatal("Head is already tagged with semver tag ", tag)
	}

}

// FindLatestSemverTag returns the latest semver tag found on current branch
// returns "","", nil if no tag can be found
func FindLatestSemverTag(repo *git.Repository) (string, plumbing.Hash, error) {
	tagList := make(map[plumbing.Hash]string)
	/* Get all tags indexed by hash */

	tags, err := repo.Tags()
	if err != nil {
		return "", plumbing.ZeroHash, err
	}

	for ref, err := tags.Next(); err == nil; ref, err = tags.Next() {
		tagName := ref.Name().Short()
		tagList[ref.Hash()] = tagName
	}
	tags.Close()

	iter, err := repo.Log(&git.LogOptions{})
	if err != nil {
		return "", plumbing.ZeroHash, err
	}
	defer iter.Close()

	for ref, err := iter.Next(); err == nil; ref, err = iter.Next() {
		tag, found := tagList[ref.Hash]
		if found {
			_, err := semver.NewVersion(tag)
			if err == nil {
				return tag, ref.Hash, nil
			}
		}
	}
	return "", plumbing.ZeroHash, nil
}

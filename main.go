package main

import (
	"fmt"
	"log"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func main() {

	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal(err)
	}

	/* Get all tags

	tags, err := repo.Tags()
	if err != nil {
		log.Fatal(err)
	}
	defer tags.Close()

	for ref, err := tags.Next(); err == nil; ref, err = tags.Next() {
		tagName := ref.Name().Short()
		log.Print(tagName)
	}
	*/

	head, err := repo.Head()
	if err != nil {
		log.Fatal(err)
	}

	refs, _ := repo.References()
	refs.ForEach(func(ref *plumbing.Reference) error {
		if ref.Type() == plumbing.HashReference && ref.Hash() == head.Hash() {
			fmt.Println(ref)
			fmt.Println(ref.Name().Short())
		}

		return nil
	})
}

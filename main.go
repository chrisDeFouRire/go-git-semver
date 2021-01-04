package main

import (
	"log"

	"github.com/go-git/go-git/v5"
)

func main() {

	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal(err)
	}

	tags, err := repo.Tags()
	if err != nil {
		log.Fatal(err)
	}
	defer tags.Close()

	for ref, err := tags.Next(); err == nil; ref, err = tags.Next() {
		tagName := ref.Name().Short()
		log.Print(tagName)
	}
}

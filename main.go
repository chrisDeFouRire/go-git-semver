package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

var (
	// ErrNotSemver if a tag can't be parsed as a semantic versioning tag
	ErrNotSemver = fmt.Errorf("Tag is not SEMVER")
	// ErrCantParseNumber if a tag number can't be parsed
	ErrCantParseNumber = fmt.Errorf("Tag number can't be parsed")

	re = regexp.MustCompile(`^v([0-9]+)\.([0-9]+)\.([0-9]+)-?(.*)$`)
)

func main() {

	tagList := make(map[plumbing.Hash]string)

	repo, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal(err)
	}

	/* Get all tags by hash */

	log.Println(">> find TAGS ")
	tags, err := repo.Tags()
	if err != nil {
		log.Fatal(err)
	}
	defer tags.Close()

	for ref, err := tags.Next(); err == nil; ref, err = tags.Next() {
		tagName := ref.Name().Short()
		tagList[ref.Hash()] = tagName
	}

	log.Println(">> Inspect log")

	iter, err := repo.Log(&git.LogOptions{})
	for ref, err := iter.Next(); err == nil; ref, err = iter.Next() {
		tag, found := tagList[ref.Hash]
		if found {
			major, minor, patch, rest, err := tag2Semver(tag)

			if err == nil {
				log.Printf("Version %d.%d.%d-%s", major, minor, patch, rest)
				break
			}
		}
	}

}

func tag2Semver(tag string) (int, int, int, string, error) {
	found := re.FindSubmatch([]byte(tag))
	if found == nil {
		return 0, 0, 0, "", ErrNotSemver
	}
	major, err := strconv.Atoi(string(found[1]))
	if err != nil {
		return 0, 0, 0, "", ErrCantParseNumber
	}
	minor, err := strconv.Atoi(string(found[2]))
	if err != nil {
		return 0, 0, 0, "", ErrCantParseNumber
	}
	patch, err := strconv.Atoi(string(found[3]))
	if err != nil {
		return 0, 0, 0, "", ErrCantParseNumber
	}

	return major, minor, patch, string(found[4]), nil
}
